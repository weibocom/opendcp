package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

//邮件的结构
type EmailData struct {
	Sender      string
	Receiver    string
	UserName    string
	Cc          string
	Subject     string
	Content     string
	MailType    string
	Attachments map[string]*Attachment
}

//发送者
type Sender struct {
	EmailName, Password, EmailServer string
}

//附件
type Attachment struct {
	Filename string
	Data     []byte
	Inline   bool
}


func (a *Sender) GetLoginName() string {
	if a.EmailName != "" {
		names := strings.Split(a.EmailName, "@")
		return names[0]
	}
	return ""
}

func (a *Sender) Start(server *smtp.ServerInfo) (string, []byte, error) {
	server.TLS = false
	return "LOGIN", []byte(a.GetLoginName()), nil
}

func (a *Sender) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.GetLoginName()), nil
		case "Password:":
			return []byte(a.Password), nil
		}
	}
	return nil, nil
}

//获取最终的[]byte
func (m *EmailData) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("From: " + m.Sender + "\r\n")

	t := time.Now()
	buf.WriteString("Date: " + t.Format("2006-01-02 15:04:05") + "\r\n")

	buf.WriteString("To: " + m.Receiver + "\r\n")

	buf.WriteString("Cc: " + m.Cc + "\r\n")

	buf.WriteString("Subject: " + m.Subject + "\r\n")

	buf.WriteString("MIME-Version: 1.0\r\n")

	boundary := "f46d043c813270fc6b04c2d223da"

	if len(m.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n")
		buf.WriteString("--" + boundary + "\r\n")
	}

	var content_type string
	if m.MailType == "html" {
		content_type = "Content-Type: text/" + m.MailType + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	buf.WriteString(fmt.Sprintf("%s\r\n\r\n", content_type))
	buf.WriteString(m.Content)
	buf.WriteString("\r\n")

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			buf.WriteString("\r\n\r\n--" + boundary + "\r\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\r\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\r\n\r\n")

				buf.Write(attachment.Data)
			} else {
				buf.WriteString("Content-Type: application/octet-stream; charset=UTF-8 \r\n")
				buf.WriteString("Content-Transfer-Encoding: base64\r\n")
				buf.WriteString("Content-Disposition: attachment; filename=\"" + attachment.Filename + "\"\r\n\r\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)

				// write base64 content in lines of up to 76 chars
				for i, l := 0, len(b); i < l; i++ {
					buf.WriteByte(b[i])
					if (i+1)%76 == 0 {
						buf.WriteString("\r\n")
					}
				}
			}

			buf.WriteString("\r\n--" + boundary)
		}
		buf.WriteString("--")
	}
	return buf.Bytes()
}

func (m *EmailData) attach(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

func (m *EmailData) Attach(file string) error {
	return m.attach(file, false)
}

func (m *EmailData) Inline(file string) error {
	return m.attach(file, true)
}
