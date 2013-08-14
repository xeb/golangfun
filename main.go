package main

import (
	"fmt"
	"github.com/xeb/golangfun/libfun"
)

func main() {
	isFun := libfun.HaveFun()
	fmt.Printf("I am main.  I am %t\n", isFun)
}