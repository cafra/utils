package utils

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

var client *mailClient

type mailClient struct {
	UserName string
	PassWord string
	Gateway  string // such as "smtp.163.com:25"
}

func SendMail(UserName, PassWord, Gateway, subject, content string, toSet []string) {
	if client == nil {
		client = &mailClient{
			UserName: UserName,
			PassWord: PassWord,
			Gateway:  Gateway,
		}
	}
	client.mail(toSet, subject, content)
}

func (m *mailClient) mail(toSet []string, subject, content string) {
	auth := sasl.NewPlainClient("", m.UserName, m.PassWord)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"chenzhen@cmcm.com"}
	msg := strings.NewReader("To: chenzhen@cmcm.com\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + content + "\r\n")
	err := smtp.SendMail(m.Gateway, auth, m.UserName, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
