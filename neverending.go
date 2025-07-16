package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			msg, ok := <-ball
			fmt.Println(msg, ok, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	ball <- "referee" // kick off
	var c chan bool   // nil
	<-c               // blocking here for ever
}
