package account

import (
	"net/smtp"
	"strings"
	. "weibo.com/opendcp/jupiter/models"
)

var (
	emailService = &EmailService{}
)

type EmailService struct {
}


//创建发送者
func (e *EmailService) NewSender(email, password, server string) *Sender {
	s := new(Sender)
	s.EmailName = email
	s.Password = password
	s.EmailServer = server
	return s
}

//创建邮件
func (e *EmailService) NewEmailDate(sender, cc, receiver, subject, content, mailType string) *EmailData {
	em := new(EmailData)
	em.Sender = sender
	em.Cc = cc
	em.Receiver = receiver
	em.Subject = subject
	em.Content = content
	em.MailType = mailType
	em.Attachments = make(map[string]*Attachment) //初始化附件
	return em
}

//发送邮件,userMail完整的邮箱
func (e *EmailService) SendMail(s *Sender, d *EmailData) error {
	auth := e.GetunencryptedAuthBySender(s)
	msg := d.Bytes()
	tos := strings.Split(d.Receiver, ";")
	ccs := strings.Split(d.Cc, ";")
	send_tos := make([]string, 0)
	for _, v := range tos {
		if v != "" {
			send_tos = append(send_tos, v)
		}
	}
	for _, v := range ccs {
		if v != "" {
			send_tos = append(send_tos, v)
		}
	}

	smtp.SendMail(s.EmailServer + ":25", auth, s.EmailName, send_tos, msg)
	return nil
}

//获取smtpAuth
func (e *EmailService) GetunencryptedAuthBySender(s *Sender) smtp.Auth {
	return &Sender{s.EmailName, s.Password, s.EmailServer}
}
