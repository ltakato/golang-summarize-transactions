package services

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"log"
	"net/textproto"
	"os"
	"summarize-transactions/config"
)

type IMAPClient struct {
	*client.Client
}

type EmailMessage struct {
	Subject        string
	SenderAddress  string
	AttachmentBody []byte
}

func NewIMAPClient(imapConfig config.ImapConfig) (*IMAPClient, error) {
	imapServer := os.Getenv("IMAP_SERVER")
	c, err := client.DialTLS(imapServer, nil)

	if err != nil {
		return nil, err
	}

	if err = c.Login(imapConfig.Email, imapConfig.Password); err != nil {
		return nil, err
	}

	return &IMAPClient{c}, nil
}

func (c *IMAPClient) SearchWithSubject(subject string) ([]uint32, error) {
	criteria := imap.SearchCriteria{
		Header: textproto.MIMEHeader{"Subject": {subject}},
	}
	seqNums, err := c.Search(&criteria)

	return seqNums, err
}

func (c *IMAPClient) FetchMessages(subject string, msgch chan *EmailMessage) error {
	seqNums, err := c.SearchWithSubject(subject)

	if err != nil {
		return err
	}

	if len(seqNums) == 0 {
		log.Printf("There's no messages to process. Finishing process")
		return nil
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNums...)

	section := &imap.BodySectionName{}

	ch := make(chan *imap.Message, 10)

	err = c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, ch)

	for m := range ch {
		senderAddress := m.Envelope.From[0].Address()
		mr, errm := mail.CreateReader(m.GetBody(section))

		if errm != nil {
			return errm
		}

		attachmentBody, errm := getMessageAttachmentBody(mr)

		if errm != nil {
			return errm
		}

		msgch <- &EmailMessage{
			Subject:        m.Envelope.Subject,
			SenderAddress:  senderAddress,
			AttachmentBody: attachmentBody,
		}
	}

	return err
}

func getMessageAttachmentBody(messageReader *mail.Reader) ([]byte, error) {
	var body []byte
	var err error

	for {
		p, errp := messageReader.NextPart()

		if errp == io.EOF {
			break
		} else if errp != nil {
			return nil, errp
		}

		switch p.Header.(type) {
		case *mail.AttachmentHeader:
			body, err = io.ReadAll(p.Body)
		}
	}

	return body, err
}
