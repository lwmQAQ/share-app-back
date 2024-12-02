package utils

import (
	"fmt"
	"login-server/config"
	"math/rand"
	"time"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type EmailUtil struct {
	config *config.EmailConfig
}

func NewEmailUtils(config *config.EmailConfig) *EmailUtil {

	return &EmailUtil{
		config: config,
	}
}

func (e *EmailUtil) sendhtmlEmail(body string, toEmail string) error {
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", e.config.From)
	// 设置收件人
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "lwm的app")
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

func (e *EmailUtil) sendTextEmail(body string, toEmail string) error {
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", e.config.From)
	// 设置收件人
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "lwm的app")
	m.SetBody("text/plain", body)
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
func (e *EmailUtil) SendHTMLEmail(name string, link string, toEmail string) error {
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

func (e *EmailUtil) SendCode(toEmail string, code string) error {
	body := fmt.Sprintf("欢迎使用我的软件，你的验证码是:%s", code)
	return e.sendTextEmail(body, toEmail)
}

func (e *EmailUtil) SendIPChangeEmail(toEmail string) {
	//TODO 处理ip异常的email
}

func CreateCode() string {
	// 创建一个新的随机数生成器，并使用当前时间作为种子
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := randGenerator.Intn(1000000) // 生成一个 0 到 999999 之间的随机数
	return fmt.Sprintf("%06d", code)    // 格式化成六位数，不足六位时前面补零
}
