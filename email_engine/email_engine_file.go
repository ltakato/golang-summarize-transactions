package email_engine

import (
	"strings"
	"summarize-transactions/services"
	"time"
)

type EmailEngineMessage struct {
	FileName       string
	AttachmentBody []byte
}

func NewEmailEngineMessage(msg *services.EmailMessage) *EmailEngineMessage {
	senderAddress := msg.SenderAddress
	formattedTime := time.Now().Format("200601021504")
	senderAddress = strings.ReplaceAll(senderAddress, ".", "_")
	fileName := senderAddress + "-" + formattedTime + ".csv"

	return &EmailEngineMessage{
		FileName:       fileName,
		AttachmentBody: msg.AttachmentBody,
	}
}
