package email

import (
	"bytes"
	"embed"
	"html/template"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"gopkg.in/gomail.v2"
)

type Credentials struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
	Token           string
}

type EmailData struct {
	// title of email
	Title string

	// sender's email address
	Sender string

	// Receiver's email address
	Receiver string

	// Email charset
	Charset string

	// Embeded template
	TemplateFs embed.FS

	// Tempalte patterns from embed like templates/*.html
	TemplatePatterns string

	// Email template name without path
	TemplateName string

	// Attachment bytes if any
	FileBytes *bytes.Buffer

	// Attachment file name. 
	// It should be with extension
	FileName string

	// Data for email body
	Input interface{}
}

func CreateAWSSession(creds Credentials) (*session.Session, error) {

	// Create AWS session using credentials
	return session.NewSession(&aws.Config{
		Region:      aws.String(creds.Region),
		Credentials: credentials.NewStaticCredentials(creds.AccessKeyId, creds.SecretAccessKey, creds.Token),
	})
}

func SendAWSSESEmail(sess *session.Session, data *EmailData) (bool, error) {

	service := ses.New(sess)

	// Get parsed email template
	rawEmail, emailTemplate, err := getEmailTemplate(data)

	// Attempt to send email
	if !rawEmail {
		_, err = service.SendEmail(emailTemplate.(*ses.SendEmailInput))
	} else {
		// Email with attachment
		_, err = service.SendRawEmail(emailTemplate.(*ses.SendRawEmailInput))
	}

	return err != nil, err
}

func getEmailTemplate(data *EmailData) (rawMail bool, emailTemplate interface{}, err error) {

	rawMail = data.FileBytes != nil

	// Get Email with data
	htmlBody, err := getEmailWithDataTemplate(data)

	if err == nil{
		if rawMail {
			emailTemplate = prepareRawEmailTemplate(htmlBody, data)
		} else {
			emailTemplate = prepareEmailTemplate(htmlBody, data)
		}	
	}

	return rawMail, emailTemplate, err
}

func getEmailWithDataTemplate(data *EmailData) (string, error) {
	// Parse email template using Embed and template patterns
	emailTemplate := template.Must(template.ParseFS(data.TemplateFs, data.TemplatePatterns))

	var templateBuffer bytes.Buffer
	err := emailTemplate.ExecuteTemplate(&templateBuffer, data.TemplateName, data.Input)

	return templateBuffer.String(), err
}

func prepareEmailTemplate(htmlBody string, data *EmailData) (template *ses.SendEmailInput) {

	template = &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(data.Receiver),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(data.Charset),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(data.Charset),
					Data:    aws.String("Email Info"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(data.Charset),
				Data:    aws.String(data.Title),
			},
		},
		Source: aws.String(data.Sender),
	}
	return
}

func prepareRawEmailTemplate(htmlBody string, data *EmailData) (template *ses.SendRawEmailInput) {

	// Prepare headers and body for raw mail
	msg := gomail.NewMessage()
	msg.SetHeader("From", data.Sender)
	msg.SetHeader("To", data.Receiver)
	msg.SetHeader("Subject", data.Title)
	msg.SetBody("text/html", htmlBody)

	// FileName should be with extension
	msg.Attach(data.FileName, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(data.FileBytes.Bytes())
		return err
	}))

	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)


	// Message for raw email
	message := ses.RawMessage{
		Data: emailRaw.Bytes(),
	}

	template = &ses.SendRawEmailInput{
		RawMessage: &message,
		Destinations: []*string{
			aws.String(data.Receiver),
		},
		Source: aws.String(data.Sender),
	}

	return template
}

