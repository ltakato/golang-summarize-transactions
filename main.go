package main

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/fs"
	"log"
	"net/textproto"
	"os"
)

func main() {
	// connect client
	imapServer := os.Getenv("IMAP_SERVER")
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")
	c, err := client.DialTLS(imapServer, nil)

	if err != nil {
		log.Fatal(err)
	}

	// defer logout
	defer func() {
		if logoutErr := c.Logout(); logoutErr != nil {
			log.Printf("Error during logout: %v", logoutErr)
		}
	}()

	// Login
	if err := c.Login(email, password); err != nil {
		log.Fatal(err)
	}

	// Select mailbox
	_, err = c.Select("Inbox", false)

	if err != nil {
		log.Fatal(err)
	}

	// Define search criteria
	nubankExtractSubject := "Extrato da fatura do Cartão Nubank"
	criteria := imap.SearchCriteria{
		Header: textproto.MIMEHeader{"Subject": {nubankExtractSubject}},
	}

	// Perform the search
	seqNums, err := c.Search(&criteria)

	done := make(chan error, 1)

	if len(seqNums) > 0 {
		// Fetch matching messages
		seqset := new(imap.SeqSet)
		seqset.AddNum(seqNums...)

		section := &imap.BodySectionName{}
		messages := make(chan *imap.Message, 10)

		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
		}()

		for msg := range messages {
			log.Println("* " + msg.Envelope.Subject)

			mr, err := mail.CreateReader(msg.GetBody(section))

			if err != nil {
				log.Fatal(err)
			}

			for {
				p, err := mr.NextPart()

				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch h := p.Header.(type) {
				case *mail.AttachmentHeader:
					filename, _ := h.Filename()

					log.Printf("Got attachment: %v", filename)

					b, errp := io.ReadAll(p.Body)

					if errp != nil {
						fmt.Println("errp ===== :", errp)
					}

					err := os.WriteFile(filename, b, fs.ModePerm)

					if err != nil {
						log.Println("attachment err: ", err)
					}
				}
			}
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		log.Println("Done!")
	}
}
