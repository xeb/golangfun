package main

import (
	"fmt"
	"time"
	"github.com/xeb/golangfun/libfun"
)

func main() {
	isFun := libfun.HaveFun()
	fmt.Printf("I am main.  I am %t\n", isFun)

	account := &libfun.Account{1, "test", time.Now()}
	fmt.Println(account.String())
}
