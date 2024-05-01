package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeSteamGroup(group_name string) []Message {
	group_url := fmt.Sprintf("https://steamcommunity.com/groups/%s/comments", group_name)
	var messages []Message

	res, err := http.Get(group_url)
	if err != nil {
		log.Printf("error fetching URL: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("error loading HTTP response body: %s", err)
	}

	doc.Find(".commentthread_comments .commentthread_comment").Each(func(i int, s *goquery.Selection) {
		message_id := strings.Split(s.AttrOr("id", ""), "_")[1]
		message_author := strings.Replace(strings.TrimSpace(s.Find("a").Text()), "@", "@\u200b", -1)
		message_text := strings.Replace(strings.TrimSpace(s.Find(".commentthread_comment_text").Text()), "@", "@\u200b", -1)

		author_page := s.Find(".commentthread_comment_avatar a").AttrOr("href", "")
		author_picture := ScrapeUserProfilePicture(author_page)

		if message_text != "" && message_text != "This comment is awaiting analysis by our automated content check system. It will be temporarily hidden until we verify that it does not contain harmful content (e.g. links to websites that attempt to steal information)." {
			messages = append(messages, Message{message_id, message_author, author_picture, message_text})
		}
	})

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages
}

func ScrapeUserProfilePicture(profile_url string) string {
	var profile_picture string

	res, err := http.Get(profile_url)
	if err != nil {
		log.Printf("error fetching URL: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("error loading HTTP response body: %s", err)
	}

	doc.Find(".playerAvatarAutoSizeInner img").Each(func(i int, s *goquery.Selection) {
		profile_picture = s.AttrOr("src", "")
	})

	return profile_picture
}
