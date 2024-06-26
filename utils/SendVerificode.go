package utils

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/k3a/html2text"
	"github.com/vanng822/go-premailer/premailer"
	"gopkg.in/gomail.v2"
)

// Config 包含邮件服务的配置信息
type Config struct {
	// 发送邮件的固定发送者邮箱地址
	EmailFrom string
	SmtpHost  string
	SmtpPort  int
	SmtpUser  string
	SmtpPass  string
}

// FormatEmail 格式化邮箱地址
func FormatEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// SendCode 只发送验证码
func SendCode(email string, code string, config Config) error {
	from := config.EmailFrom
	smtpPass := config.SmtpPass
	smtpUser := config.SmtpUser
	to := FormatEmail(email)
	smtpHost := config.SmtpHost
	smtpPort := config.SmtpPort

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/html", code)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return errors.New("could not send email")
	}
	return nil
}

// SendVerificationCode 发送验证码至指定邮箱
func SendVerificationCode(email string, code string, config Config) error {
	from := config.EmailFrom
	smtpPass := config.SmtpPass
	smtpUser := config.SmtpUser
	to := FormatEmail(email)
	smtpHost := config.SmtpHost
	smtpPort := config.SmtpPort

	var body bytes.Buffer
	templateData := struct {
		Code string
	}{
		Code: code,
	}

	var paths []string
	err := filepath.Walk("./utils/SendVerificode", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		fmt.Print("could not read template directory")
	}

	template, err := template.ParseFiles(paths...)
	if err != nil {
		fmt.Print("could not parse template")
	}

	template.ExecuteTemplate(&body, "email-temp.html", &templateData)
	htmlString := body.String()
	prem, _ := premailer.NewPremailerFromString(htmlString, nil)
	htmlInline, err := prem.Transform()
	if err != nil {
		fmt.Print("could not inline css")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/html", htmlInline)
	m.AddAlternative("text/plain", html2text.HTML2Text(htmlString))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return errors.New("could not send email")
	}
	return nil
}
