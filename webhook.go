package main

import (
	"fmt"
	"log"

	"github.com/gtuk/discordwebhook"
)

func PostMessage(content Message, url string, show_steam_prefix bool) {
	username_tag := content.message_author
	if show_steam_prefix {
		username_tag = fmt.Sprintf("[Steam] %s", content.message_author)
	}

	message := discordwebhook.Message{
		Username:  &username_tag,
		AvatarUrl: &content.author_picture,
		Content:   &content.message_text,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Printf("error sending message: %s", err)
	}
}
