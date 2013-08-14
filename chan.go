package main

import (
	"fmt"
	"time"
)

func FixedMessagePump() {
	
	fmt.Println("Starting...")
	messagePump := make(chan string)

	limit := 145

	for i :=0; i < limit; i++ {
		go func(j int) { messagePump <- fmt.Sprintf("ping-%d", j) }(i)
	}

	for i :=0; i < limit; i++ {
		lastMsg := <-messagePump
		fmt.Println(lastMsg)
	}

	fmt.Println("Done!")
}

func TimeoutExample() {
	expectChan := make(chan bool, 1)

	// Standard operation that takes 10 seconds
	go func() { 
		time.Sleep(5 * time.Second)
		expectChan <- true
	}()

	select {
		case <-expectChan:
			fmt.Println("Done in expected time from chan")
		case <-time.After(3 * time.Second):
			fmt.Println("Timed out")
	}
}

func main() {

	TimeoutExample();
	// FixedMessagePump();
}