package email

import (
	"embed"
	"github.com/canopas/go-scaffolds/file"
	"html/template"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type emailInput struct {
	Name  string
	Email string
}

//go:embed templates/*.html
var templateFS embed.FS

var parsedTemplate = template.Must(template.ParseFS(templateFS, "templates/*.html"))
var testEmail = "test@gmail.com"

var dataInput = emailInput{
	Name:  "Test",
	Email: testEmail,
}

var emailData = EmailData{
	Title:            "Test email",
	Subject:          "Test email subject",
	Sender:           testEmail,
	Receiver:         testEmail,
	Charset:          "UTF-8",
	TemplateFs:       templateFS,
	TemplatePatterns: "templates/*.html",
	Input:            dataInput,
}

func TestGetHTMLOfEmailTemplateError(t *testing.T) {
	asserts := assert.New(t)
	temp, err := getEmailWithDataTemplate(&emailData)
	asserts.Error(err)
	asserts.Empty(temp)
}

func TestGetHTMLOfEmailTemplateSuccess(t *testing.T) {
	asserts := assert.New(t)
	emailData.TemplateName = "test.html"
	temp, err := getEmailWithDataTemplate(&emailData)
	asserts.NoError(err)
	asserts.NotEmpty(temp)
}

func TestPrepareEmailTemplateSuccess(t *testing.T) {
	asserts := assert.New(t)
	temp := prepareEmailTemplate("", &emailData)
	asserts.NotEmpty(temp)
}

func TestPrepareRawEmailTemplateSuccess(t *testing.T) {
	asserts := assert.New(t)
	prepareEmailData(true)
	temp := prepareRawEmailTemplate("", &emailData)
	asserts.NotEmpty(temp)
}

func TestGetEmailTemplateRawSuccess(t *testing.T) {
	asserts := assert.New(t)
	prepareEmailData(true)
	rawMail, template, err := getEmailTemplate(&emailData)

	asserts.Nil(err)
	asserts.NotEmpty(template)
	asserts.Equal(true, rawMail)
}

func TestGetEmailTemplateSuccess(t *testing.T) {
	asserts := assert.New(t)
	prepareEmailData(false)

	rawMail, template, err := getEmailTemplate(&emailData)

	asserts.Nil(err)
	asserts.NotEmpty(template)
	asserts.Equal(false, rawMail)
}

func prepareEmailData(rawMail bool) {
	emailData.TemplateFs = templateFS
	emailData.TemplatePatterns = "templates/*.html"
	emailData.TemplateName = "test.html"

	if !rawMail {
		emailData.FileBytes = nil
	} else {
		fileBytes, err := file.GetFileBytes("./test.txt")

		if err != nil {
			log.Error("err:", err)
		}

		emailData.FileBytes = fileBytes
		emailData.FileName = "test.txt"
	}
}
