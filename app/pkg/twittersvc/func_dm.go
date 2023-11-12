package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
	"log"
)

func (s *TwitterServiceMaster) DirectMessage(ctx core.Context) (err error) {

	// 发送私信
	sendDirectMessage(s.xClient, "recipientScreenName", "Hello, this is a direct message.")
	return nil
}

// 发送私信
func sendDirectMessage(client *twitter.Client, recipientScreenName, message string) {
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
}
