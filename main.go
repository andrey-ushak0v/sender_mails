package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net/smtp"
)

type UserInDb struct {
	FirstName string
	LastName  string
	BirthYear int
}

type TemplateUser struct {
	Name    string
	Surname string
	Born    int
}

func GetUserByEmail(mail string, users map[string]UserInDb) UserInDb {
	return users[mail]
}

func SendMail(subject string, templatePath string, to string, user UserInDb) {

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
		"uandrey220@gmail.com",
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	users := map[string]UserInDb{

		"foo@baiirhhh.com": UserInDb{FirstName: "Vaspp", LastName: "Pkmnin", BirthYear: 1981},
	}

	usersEmail := flag.String("email", "example@gmail.com", "a string")
	flag.Parse()
	SendMail(
		"very important message",
		"./maket.html",
		*usersEmail,
		GetUserByEmail(*usersEmail, users),
	)
	fmt.Println(*usersEmail)
}
