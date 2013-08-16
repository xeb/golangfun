package main

import (
	"fmt"
	"time"
	"github.com/xeb/golangfun/libfun"
	"github.com/xeb/golangfun/channels"
)

func main() {

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
