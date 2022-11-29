package main

import (
	"flag"
	"fmt"
	"sender_mails/models"
	"sender_mails/send_meaasge"
)

func main() {

	usersEmail := flag.String("email", "example@gmail.com", "a string")
	flag.Parse()
	send_meaasge.SendMail(
		"very important message",
		"./maket.html",
		*usersEmail,
		models.GetUserByEmail(*usersEmail, models.UsersData()),
	)
	fmt.Println(*usersEmail)
}
