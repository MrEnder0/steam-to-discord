package main

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

func PostMessage(content Message, url string) {
	message := discordwebhook.Message{
		Username: &content.message_author,
		Content:  &content.message_text,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Printf("error sending message: %s", err)
	}
}
