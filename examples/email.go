package examples

import (
	"embed"

	"github.com/canopas/go-scaffolds/email"
	"github.com/canopas/go-scaffolds/file"

	log "github.com/sirupsen/logrus"
)

type MailData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//go:embed templates/*.html
var templateFS embed.FS

func SendEmail() {

	// if AWS sess in not created before, create it using CreateAWSSession or skip this
	sess, err := email.CreateAWSSession(email.Credentials{
		Region:          "YOUR-REGION",
		AccessKeyId:     "YOUR-ACCESS-KEY-ID",
		SecretAccessKey: "YOUR-SECRET-ACCESS-KEY",
		Token:           "",
	})

	if err != nil {
		log.Error("err:", err)
	}

	// Prepare your email data which you want in your email template
	inputData := MailData{
		Name:  "name",
		Email: "email",
	}

	data := &email.EmailData{
		Title:            "email-title",
		Subject:          "email-subject",
		Sender:           "sender-email",
		Receiver:         "receiver-email",
		Charset:          "UTF-8",
		TemplateFs:       templateFS,
		TemplatePatterns: "templates/*.html",
		TemplateName:     "example.html",
		Input:            inputData,
	}

	_, err = email.SendAWSSESEmail(sess, data)

	if err != nil {
		log.Error("err:", err)
	}
}

func SendEmailWithAttachment() {

	// if AWS sess in not created, create it using CreateAWSSession or skip this
	sess, err := email.CreateAWSSession(email.Credentials{
		Region:          "YOUR-REGION",
		AccessKeyId:     "YOUR-ACCESS-KEY-ID",
		SecretAccessKey: "YOUR-SECRET-ACCESS-KEY",
		Token:           "",
	})

	if err != nil {
		log.Error("err:", err)
	}

	// Prepare your email data which you want in your email template
	inputData := MailData{
		Name:  "name",
		Email: "email",
	}

	// If email has an attacment, then create filebytes
	fileBytes, err := file.GetFileBytes("./test.txt")

	if err != nil {
		log.Error("err:", err)
	}

	data := &email.EmailData{
		Title:            "email-title",
		Sender:           "sender-email",
		Receiver:         "receiver-email",
		Charset:          "UTF-8",
		TemplateFs:       templateFS,
		TemplatePatterns: "templates/*.html",
		TemplateName:     "example.html",
		FileBytes:        fileBytes,
		FileName:         "test.txt",
		Input:            inputData,
	}

	_, err = email.SendAWSSESEmail(sess, data)

	if err != nil {
		log.Error("err:", err)
	}
}
