package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/xeb/golangfun/libfun"
	"github.com/xeb/golangfun/channels"
)


var flagparam bool

func init() {
	flag.BoolVar(&flagparam, "flagparam", false, "Flag Parameter")
	flag.Parse()
}

func main() {

	if flagparam {
		fmt.Println("Flagparam is set")
	}

	libfun.TestPointer()
	libfun.Hello()

	isFun := libfun.HaveFun()
	fmt.Printf("I am main.  I am %t\n", isFun)

	i := libfun.DeferFun()
	fmt.Printf("Value of i is %d\n", i)

	account := &libfun.Account{1, "test", time.Now()}
	fmt.Println(account.String())

	channels.TimeoutExample()
	// channels.FixedMessagePump()
}
