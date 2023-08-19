package mail

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

// Sender is an email sender
type Sender struct {
	fromName     string
	fromEmail    string
	fromPassword string
}

// SendEmail sends an email to the given recipients
func (s *Sender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.fromName, s.fromEmail)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.Text = []byte(content)
	for _, attachFile := range attachFiles {
		_, err := e.AttachFile(attachFile)
		if err != nil {
			return errors.Join(
				err,
				fmt.Errorf("failed to attach file %s", attachFile),
			)
		}
	}

	err := e.Send(
		smtpAuthAddress,
		smtp.PlainAuth(
			"",
			s.fromEmail,
			s.fromPassword,
			smtpServerAddress,
		),
	)

	if err != nil {
		return errors.Join(
			err,
			fmt.Errorf("failed to send email"),
		)
	}

	return nil
}

// NewEmailSender creates a new email sender
func NewEmailSender(fromName, fromEmail, fromPassword string) IEmailSender {
	return &Sender{
		fromName:     fromName,
		fromEmail:    fromEmail,
		fromPassword: fromPassword,
	}
}
