package send_meaasge

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"sender_mails/models"
)

type TemplateUser struct {
	Name    string
	Surname string
	Born    int
}

func SendMail(subject string, templatePath string, to string, user models.UserInDb) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, TemplateUser{Name: user.FirstName, Surname: user.LastName, Born: user.BirthYear})

	auth := smtp.PlainAuth(
		"",
		"",
		"",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"",
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
