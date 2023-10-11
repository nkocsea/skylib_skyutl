package skyutl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/nkocsea/skylib_skylog/skylog"
)

type ReadDocInfo struct {
	CompanyId   int64  `json:"companyId"`
	BranchId    int64  `json:"branchId"`
	Iuid        string `json:"iuid"`
	ServiceCode string `json:"serviceCode"`
	ScreenCode  string `json:"screenCode"`
	FeatureCode string `json:"featureCode"`
	FullPath    string `json:"fullPath"`
}

type WriteDocInfo struct {
	CompanyId   int64  `json:"companyId"`
	BranchId    int64  `json:"branchId"`
	Iuid        string `json:"iuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Note        string `json:"note"`
	ServiceCode string `json:"serviceCode"`
	ScreenCode  string `json:"screenCode"`
	FeatureCode string `json:"featureCode"`
	ItemType    string `json:"itemType"`
	ItemCode    string `json:"itemCode"`
	ItemId      int64  `json:"itemId"`
	ItemDate    int64  `json:"itemDate"`
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	Mode        int32  `json:"mode"`
	Md5         string `json:"md5"`
	Size        int64  `json:"size"`
	Force       bool   `json:"force"`
}

func WriteDoc(ctx context.Context, urlOrServiceAddress string, metadata WriteDocInfo, data []byte) ([]byte, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	fileData, err := writer.CreateFormFile("data", metadata.Name)
	if err != nil {
		skylog.Errorf("WriteBusinessFile create form file error: %v", err)
		return nil, err
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(data)
	_, err = io.Copy(fileData, reader)
	if err != nil {
		skylog.Errorf("WriteBusinessFile copy byte array error error: %v", err)
		return nil, err
	}
	_ = writer.WriteField("companyId", ToStr(metadata.CompanyId))
	_ = writer.WriteField("branchId", ToStr(metadata.BranchId))
	_ = writer.WriteField("iuid", metadata.Iuid)
	_ = writer.WriteField("name", metadata.Name)
	_ = writer.WriteField("description", metadata.Description)
	_ = writer.WriteField("note", metadata.Note)
	_ = writer.WriteField("serviceCode", metadata.ServiceCode)
	_ = writer.WriteField("screenCode", metadata.ScreenCode)
	_ = writer.WriteField("featureCode", metadata.FeatureCode)
	_ = writer.WriteField("itemType", metadata.ItemType)
	_ = writer.WriteField("itemCode", metadata.ItemCode)
	_ = writer.WriteField("itemId", ToStr(metadata.ItemId))
	_ = writer.WriteField("itemDate", ToStr(metadata.ItemDate))
	_ = writer.WriteField("partnerCode", metadata.PartnerCode)
	_ = writer.WriteField("partnerName", metadata.PartnerName)
	_ = writer.WriteField("mode", ToStr(metadata.Mode))
	_ = writer.WriteField("md5", metadata.Md5)
	_ = writer.WriteField("size", ToStr(metadata.Size))
	_ = writer.WriteField("force", ToStr(metadata.Force))

	err = writer.Close()
	if err != nil {
		skylog.Errorf("WriteBusinessFile close write error: %v", err)
		return nil, err
	}

	accessToken, _, _ := GetLoginAccessToken(ctx)
	urlPath := "/doc/file/v1/upload"
	return RestUploadFile(urlOrServiceAddress, urlPath, payload, writer.FormDataContentType(), accessToken)
}

func ReadDocWithUrl(ctx context.Context, fullUrl string) ([]byte, string, string, string, error) {
	accessToken, _, _ := GetLoginAccessToken(ctx)
	return RestDownloadFile(fullUrl, "", accessToken)
}

func ReadDoc(ctx context.Context, urlOrServiceAddress string, metadata ReadDocInfo) ([]byte, string, string, string, error) {
	accessToken, _, _ := GetLoginAccessToken(ctx)

	urlPath := ""
	fullPath := strings.TrimSpace(metadata.FullPath)
	if len(fullPath) > 0 {
		urlPath = BuildDocUrlWithFullPath(fullPath)
	} else {
		urlPath = BuildDocUrl(metadata.Iuid, metadata.CompanyId, metadata.BranchId, metadata.ServiceCode, metadata.ScreenCode, metadata.FeatureCode)
	}
	return RestDownloadFile(urlOrServiceAddress, urlPath, accessToken)
}

func BuildDocUrlWithData(params url.Values) string {
	returnVal := fmt.Sprintf("/doc/file/v1/download?%v", params.Encode())
	skylog.Info("BuildDocUrl url: %v", returnVal)
	return returnVal
}

func BuildDocUrlWithFullPath(fullPath string) string {
	params := url.Values{}
	params.Add("fullPath", fullPath)

	return BuildDocUrlWithData(params)
}

func BuildDocUrl(iuid string, companyId, branchId int64, serviceCode, screenCode, featureCode string) string {
	params := url.Values{}
	params.Add("iuid", iuid)
	params.Add("companyId", ToStr(companyId, "0"))
	params.Add("branchId", ToStr(branchId, "0"))
	params.Add("service", serviceCode)
	params.Add("screen", screenCode)
	params.Add("feature", featureCode)

	return BuildDocUrlWithData(params)
}

func BuildDocUrlWithToken(documentId int64, token string) string {
	params := url.Values{}
	params.Add("id", ToStr(documentId, ""))
	params.Add("token", token)

	return BuildDocUrlWithData(params)
}
