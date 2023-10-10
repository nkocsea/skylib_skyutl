package skyutl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"suntech.com.vn/skylib/skylog.git/skylog"
)

func BuildServiceUrl(urlOrServiceAddr, path string) (string, error) {
	if strings.HasPrefix(urlOrServiceAddr, "http") {
		return urlOrServiceAddr, nil
	}

	parts := strings.Split(urlOrServiceAddr, ":")
	if len(parts) < 2 {
		skylog.Errorf("service address is incorrect: %v", urlOrServiceAddr)
		return "", errors.New("service address is incorrect")
	}
	grpcPort, _ := strconv.Atoi(parts[1])

	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%v", path)
	}
	url := fmt.Sprintf("http://%v:%v%v", parts[0], grpcPort+1, path)

	return url, nil
}

func SendRawRequest(method, urlOrServiceAddr, path, accessToken string, body map[string]interface{}) (*http.Response, error) {
	// build url
	serviceUrl, err := BuildServiceUrl(urlOrServiceAddr, path)
	if err != nil {
		skylog.Errorf("SendRawRequest build service url error: %v", err)
		return nil, err
	}

	// data
	var data io.Reader = nil
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		bodyData, err := json.Marshal(body)
		if err != nil {
			skylog.Errorf("SendRawRequest create body data error: %v | data: %v", err, body)
			return nil, err
		}
		data = bytes.NewReader(bodyData)
	}
	// create request
	req, err := http.NewRequest(method, serviceUrl, data)
	if err != nil {
		skylog.Errorf("SendRawRequest create request error: %v", err)
		return nil, err
	}

	// set request param
	if len(accessToken) > 0 {
		if !strings.HasPrefix(accessToken, "Bearer ") {
			accessToken = fmt.Sprintf("Bearer %v", accessToken)
		}
		req.Header.Set("Authorization", accessToken)
	}
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}

	// create http client
	client := http.Client{}
	// send request
	return client.Do(req)
}

func SendMultiPartForm(method, urlOrServiceAddr, path string, body io.Reader, contentType, accessToken string) (*http.Response, error) {
	// build url
	serviceUrl, err := BuildServiceUrl(urlOrServiceAddr, path)
	if err != nil {
		skylog.Errorf("SendMultiPartForm build service url error: %v", err)
		return nil, err
	}

	// create request
	req, err := http.NewRequest(method, serviceUrl, body)
	if err != nil {
		skylog.Errorf("SendMultiPartForm create request error: %v", err)
		return nil, err
	}

	// set request param
	if len(accessToken) > 0 {
		if !strings.HasPrefix(accessToken, "Bearer ") {
			accessToken = fmt.Sprintf("Bearer %v", accessToken)
		}
		req.Header.Set("Authorization", accessToken)
	}
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", contentType)
	}

	// create http client
	client := http.Client{}
	// send request
	return client.Do(req)
}

func SendRequest(method, urlOrServiceAddr, path, accessToken string, body map[string]interface{}) ([]byte, error) {
	// send request
	resp, err := SendRawRequest(method, urlOrServiceAddr, path, accessToken, body)
	if err != nil {
		skylog.Errorf("SendRequest send request error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	resData, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		return resData, nil
	}
	skylog.Errorf("SendRequest send request error: %v", string(resData))
	return resData, CallRESTAPIError
}

func RestGet(urlOrServiceAddr, path string, accessToken string) ([]byte, error) {
	return SendRequest(http.MethodGet, urlOrServiceAddr, path, accessToken, nil)
}

func RestPost(urlOrServiceAddr, path string, values map[string]interface{}, accessToken string) ([]byte, error) {
	return SendRequest(http.MethodPost, urlOrServiceAddr, path, accessToken, values)
}

func RestDownloadFile(urlOrServiceAddr, path string, accessToken string) ([]byte, string, string, string, error) {
	// send request
	resp, err := SendRawRequest(http.MethodGet, urlOrServiceAddr, path, accessToken, nil)
	if err != nil {
		skylog.Errorf("RestDownloadFile send request error: %v", err)
		return nil, "", "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Get the file name from the Content-Disposition header
		fileName := resp.Header.Get("Content-Disposition")
		fileName = strings.TrimSpace(strings.Split(fileName, ";")[1])
		fileName = strings.TrimSpace(strings.Split(fileName, "=")[1])
		fileName = strings.Replace(fileName, "\"", "", -1)

		// Get the content type from the Content-Type header
		contentType := resp.Header.Get("Content-Type")

		// Get the date from the Last-Modified header
		dateStr := resp.Header.Get("Last-Modified")
		if dateStr != "" {
			date, err := time.Parse(http.TimeFormat, dateStr)
			if err != nil {
				skylog.Errorf("RestDownloadFile parse time error; time=%v; err:%v", dateStr, err)
				dateStr = ""
			} else {
				dateStr = ToStr(date.UTC().UnixMilli())
			}
		}

		// Read the response body into a byte array
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			skylog.Errorf("RestDownloadFile read body error: %v", err)
			return nil, fileName, contentType, dateStr, err
		}

		return bytes, fileName, contentType, dateStr, nil
	} else {
		skylog.Errorf("RestDownloadFile call request error: %v", resp.StatusCode)
		// Read the response body into a byte array
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			skylog.Errorf("RestDownloadFile read body error: %v", err)
			return nil, "", "", "", err
		}

		return bytes, "", "", "", errors.New(string(bytes))
	}
}

func RestUploadFile(urlOrServiceAddr, path string, body io.Reader, contentType, accessToken string) ([]byte, error) {
	// send request
	resp, err := SendMultiPartForm(http.MethodPost, urlOrServiceAddr, path, body, contentType, accessToken)
	if err != nil {
		skylog.Errorf("RestUploadFile send multi part file error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	resData, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		return resData, nil
	}
	skylog.Errorf("RestUploadFile send multi part file error: %v", string(resData))
	return resData, CallRESTAPIError
}
