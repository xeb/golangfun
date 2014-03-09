package channels

import (
	"fmt"
	"time"
)

func SimpleChanValue() (value int) {
	channel := make(chan int)

	go func() {
		channel <- 1337
	}()

	for value = <-channel; value != 1337; {
		fmt.Print("\t\ttick")
	}

	fmt.Println("Done!  Value is %d", value)

	return
}

func FixedMessagePump() {
	fmt.Println("Starting...")
	messagePump := make(chan string)

	limit := 145

	for i := 0; i < limit; i++ {
		go func(j int) { messagePump <- fmt.Sprintf("ping-%d", j) }(i)
	}

	for i := 0; i < limit; i++ {
		lastMsg := <-messagePump
		fmt.Println(lastMsg)
	}

	fmt.Println("Done!")
}

func TimeoutExample() {
	expectChan := make(chan bool, 1)

	// Standard operation that takes 10 seconds
	go func() {
		time.Sleep(2 * time.Second)
		expectChan <- true
	}()

	select {
	case <-expectChan:
		fmt.Println("Done in expected time from chan")
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out")
	}
}

func ReadWriteExample() {
	ch := make(chan int)

	reader := readChan(ch)
	writer := writeChan(ch)

	fmt.Printf("Channels created\n")
	go func() {
		val := <-reader
		fmt.Printf("Read %s\n", val)
	}()

	fmt.Println("Writing 14")
	writer <- 14
	//reader <- 14 /* will not compile */
	fmt.Println("Wrote 14")
}

func readChan(ch chan int) <-chan int {
	return ch
}

func writeChan(ch chan int) chan<- int {
	return ch
}
