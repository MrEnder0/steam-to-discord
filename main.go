package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Steam -> Discord Bridge")

	config := LoadConfig()

	CreateDatabase()

	for {
		messages := ScrapeSteamGroup(config.GroupName)
		for _, message := range messages {
			if FindMessage(message.message_id) {
				continue
			}
			PostMessage(message, config.WebhookURL, config.ShowSteamPrefix)
			InsertMessage(message.message_id)
			fmt.Printf("Posted message from: %s\n", message.message_author)
			time.Sleep(1 * time.Second)
		}
		time.Sleep(time.Duration(config.CheckFreq))
	}
}
