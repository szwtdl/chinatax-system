package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func SendEmailTLS(to, subject, body, smtpHost, username, password string, smtpPort int) error {
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	msg := []byte(fmt.Sprintf(
		`From: "智能财务平台" <%s>
To: %s
Subject: %s
MIME-Version: 1.0
Content-Type: text/html; charset=UTF-8

%s`, username, to, subject, body))

	// 建立 TLS 连接
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}
	defer c.Quit()

	auth := smtp.PlainAuth("", username, password, smtpHost)
	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(username); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = wc.Write(msg)
	return err
}
