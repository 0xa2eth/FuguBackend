package twittersvc

import (
	"log"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) DirectMessage(
	ctx core.Context,
	recipientScreenName,
	message string) (err error) {

	// 发送私信
	return sendDirectMessage(s.xClient, recipientScreenName, message)

}

// sendDirectMessage 发送私信
func sendDirectMessage(client *twitter.Client,
	recipientScreenName,
	message string) error {
	_, _, err := client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type: "message_create",
			Message: &twitter.DirectMessageEventMessage{
				SenderID: "",
				Target:   &twitter.DirectMessageTarget{RecipientID: recipientScreenName},
				Data: &twitter.DirectMessageData{
					Text:       message,
					Entities:   nil,
					Attachment: nil,
					QuickReply: nil,
					CTAs:       nil,
				},
			},
		},
	})
	if err != nil {
		log.Println("Failed to send direct message:", err)

	} else {
		log.Println("Direct message sent to", recipientScreenName)
	}
	return err
}
