package main

import (
	"time"
)

func main() {
	config := LoadConfig()
	db := OpenDatabase()

	for {
		messages := ScrapeSteamGroup(config.GroupName)
		for _, message := range messages {
			if FindMessage(db, message.message_id) {
				continue
			}
			PostMessage(message, config.WebhookURL)
			InsertMessage(db, message)
			time.Sleep(1 * time.Second)
		}
		time.Sleep(time.Duration(config.CheckFreq))
	}
}
