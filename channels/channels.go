package channels

import (
	"fmt"
)

func ChanTest() (value int) {
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

