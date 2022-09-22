package emailer

import (
	"github.com/civet148/log"
	"testing"
)

func TestSendEmail(t *testing.T) {

	m := NewEmailSender(&Config{
		SmtpServer:   "smtp.gmail.com",
		SmtpPort:     465,
		EmailAddress: "your@gmail.com",
		AuthCode:     "youremailpassword",
		EmailName:    "",
	})
	err := m.SendMail("hello", "someone@gmail.com", nil, "hi ~ I'm a test message")
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("send ok")
	}
}
