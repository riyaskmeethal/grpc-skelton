package mail

import (
	"fmt"
	"net/smtp"
)

const (
	smtpServer   = "mailv.emirates.net.ae" //smtp.gmail.com is used for gmail accounts
	smtpPort     = 25
	from         = "benefitbeyond@benefitbeyond.com"
	smtpPassword = "Benefit@123#"
	smtpUsername = "benefitbeyond@benefitbeyond.com"
)

func SendMail(to, subject, body string) error {
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body)
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	return smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
}
