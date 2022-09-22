package emailer

import (
	"github.com/civet148/log"
	"gopkg.in/gomail.v2"
)

type Config struct {
	SmtpServer   string `json:"smtp_server"`
	SmtpPort     int    `json:"smtp_port"`
	EmailAddress string `json:"email_address"`
	AuthCode     string `json:"auth_code"`
	EmailName    string `json:"email_name"`
}

type EmailSender struct {
	cfg *Config
}

func NewEmailSender(cfg *Config) *EmailSender {
	return &EmailSender{
		cfg: cfg,
	}
}

func (s *EmailSender) SendMail(subject, to string, cc []string, message string, images ...string) (err error) {
	m := gomail.NewMessage()
	//发送人
	m.SetAddressHeader("From", s.cfg.EmailAddress, s.cfg.EmailName)
	//接收人
	m.SetHeader("To", to)
	//添加抄送
	for _, c := range cc {
		m.SetAddressHeader("Cc", c, "")
	}
	//主题
	m.SetHeader("Subject", subject)
	// 内嵌图片
	for _, im := range images {
		m.Embed(im)
	}
	//内容
	m.SetBody("text/html", message)

	d := gomail.NewDialer(s.cfg.SmtpServer, s.cfg.SmtpPort, s.cfg.EmailAddress, s.cfg.AuthCode)

	// 发送邮件
	if err = d.DialAndSend(m); err != nil {
		log.Errorf("Send mail failed,err: %s", err.Error())
		return err
	}
	return nil
}
