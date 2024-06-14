package main

type Message struct {
	message_id     int
	message_author string
	author_picture string
	message_text   string
}

type Config struct {
	GroupName       string
	WebhookURL      string
	CheckFreq       int
	ShowSteamPrefix bool
}
