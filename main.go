package main

import (
	"fmt"
	"time"
)

func emit(stream chan string) {
	data := []string{"a", "b", "c", "pferd", "ochse"}
	for _, event := range data {
		stream <- event
		time.Sleep(2 * time.Second)
	}
}

func subscribe(subscriberID string, stream chan string) {
	for event := range stream {
		fmt.Println(subscriberID, event)
	}
}

func main() {
	stream := make(chan string)
	go emit(stream)

	go subscribe("sub1", stream)
	go subscribe("sub2", stream)
	go subscribe("sub3", stream)
	go subscribe("sub4", stream)
	subscribe("sub5", stream)
}
