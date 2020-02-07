package service

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapService struct {
}

func (s ImapService) List(url string, username string, password string) []string {
	log.Println("Connecting to server...")
	c, err := client.DialTLS(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	defer c.Logout()

	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	list := make([]string, len(mailboxes))
	for m := range mailboxes {
		log.Println("* " + m.Name)
		list = append(list, m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	return list
}
