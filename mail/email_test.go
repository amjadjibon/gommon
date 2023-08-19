package mail

import (
	"os"
	"testing"
)

func TestNewEmailSender(t *testing.T) {
	sender := NewEmailSender("fromName", "fromEmail", "fromPassword")
	if sender == nil {
		t.Errorf("NewEmailSender should not return nil")
	}
}

func TestSender_SendEmail(t *testing.T) {
	senderName := os.Getenv("SENDER_NAME")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	sender := NewEmailSender(senderName, senderEmail, senderPassword)

	subject := "Test Subject"
	content :=
		`
			<h1> Test Content </h1>
			<p> Test Content </p>
		`

	err := sender.SendEmail(
		subject,
		content,
		[]string{"test@email.com"},
		[]string{},
		[]string{},
		[]string{},
	)

	if err != nil {
		t.Errorf("SendEmail should not return error")
	}
}
