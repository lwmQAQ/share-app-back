package utils

import (
	"fmt"
	"login-server/config"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type EmailUtils struct {
	config *config.EmailConfig
}

func NewEmailUtils(config *config.EmailConfig) *EmailUtils {

	return &EmailUtils{
		config: config,
	}
}

func (e *EmailUtils) sendhtmlEmail(body string, toEmail string) error {
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", e.config.From)
	// 设置收件人
	m.SetHeader("To", toEmail)
	m.SetBody("text/html", body)
	// 设置SMTP服务器信息
	d := gomail.NewDialer(e.config.Host, e.config.Port, e.config.From, e.config.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送邮件时出错:", err)
		return err
	}

	fmt.Println("邮件发送成功！")
	return nil
}

/*
发送修改密码
*/
func (e *EmailUtils) SendHTMLEmail(name string, link string, toEmail string) error {
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  link,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	return e.sendhtmlEmail(emailBody, toEmail)
}
