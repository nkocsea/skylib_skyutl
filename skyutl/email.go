package skyutl

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strconv"
	"strings"
)

type Mail struct {
	Sender      string
	To          []string
	Cc          []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func SendMail(from string, password string, to, cc []string, smtpHost string, smtpPort int32, subject string, body string, attachments []string) error {
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	addr := smtpHost + ":" + strconv.Itoa(int(smtpPort))
	if len(attachments) > 0 {
		request := Mail{
			Sender:      from,
			To:          to,
			Cc:          cc,
			Subject:     subject,
			Body:        body,
			Attachments: make(map[string][]byte),
		}
		for _, fullFilePath := range attachments {
			request.attachFile(fullFilePath)
		}
		msg := BuildMailWithAttachment(request)
		err := smtp.SendMail(addr, auth, from, to, []byte(msg))
		if err != nil {
			return err
		}
	} else {
		request := Mail{
			Sender:  from,
			To:      to,
			Cc:      cc,
			Subject: subject,
			Body:    body,
		}
		msg := BuildMessage(request)
		err := smtp.SendMail(addr, auth, from, to, []byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}

func BuildMessage(mail Mail) string {
	msg := ""
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
	}

	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func BuildMailWithAttachment(mail Mail) []byte {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))

	//attach a text file to email
	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	//A multipart/mixed MIME message is composed of a mix of different data types. Each body part is delineated by a boundary. The boundary parameter is a text string used to delineate one part of the message body from another.
	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))

	for k, data := range mail.Attachments {
		//Here we define the body part, which is plain text.
		buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
		buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
		buf.WriteString("Content-Transfer-Encoding: base64\r\n")
		buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))
		buf.WriteString("Content-ID: <SA22110024.pdf>\r\n\r\n")

		//We read the data from the file.
		b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(b, data)
		buf.Write(b)
		buf.WriteString(fmt.Sprintf("\n--%s", boundary))
	}

	//We write the base64 encoded data into the buffer. The last boundary is ended with two dash characters.
	buf.WriteString("--")

	return buf.Bytes()
}

func (m *Mail) attachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}
