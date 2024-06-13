package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenDatabase() []*int {
	fi, err := os.Open("arm_database.dat")
	if err != nil {
		panic(err)
	}

	defer fi.Close()
	r := bufio.NewReader(fi)
	var messages []*int

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		lineStr := string(line)
		lineInt, err := strconv.Atoi(lineStr)
		if err != nil {
			log.Printf("Error converting line to int: %v\n", err)
			continue
		}
		messages = append(messages, &lineInt)
	}

	return messages
}

func CreateDatabase() {
	_, err := os.Stat("arm_database.dat")
	if err == nil {
		return
	}

	fi, err := os.Create("arm_database.dat")
	if err != nil {
		panic(err)
	}

	defer fi.Close()
}

func PopulateDatabase() {
	messages := OpenDatabase()
	for _, message := range messages {
		fmt.Printf("Message: %d\n", *message)
	}
}

func InsertMessage(message_id int) {
	db := OpenDatabase()
	db = append(db, &message_id)
	fi, err := os.Create("arm_database.dat")
	if err != nil {
		log.Printf("Error creating database file: %v\n", err)
		return
	}

	defer fi.Close()
	w := bufio.NewWriter(fi)
	for _, message := range db {
		fmt.Fprintln(w, *message)
	}

	w.Flush()
}

func FindMessage(message_id int) bool {
	db := OpenDatabase()
	for _, message := range db {
		if *message == message_id {
			return true
		}
	}
	return false
}
