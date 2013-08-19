package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/xeb/golangfun/libfun"
	"github.com/xeb/golangfun/lrucache"
	"github.com/xeb/golangfun/channels"
)

type Account struct {
	id int
	name string
}

var cacheEnabled bool
var channelsEnabled bool
var libfunEnabled bool

func init() {
	flag.BoolVar(&cacheEnabled, "cache", true, "Run a LRU Cache sample")
	flag.BoolVar(&channelsEnabled, "channels", false, "Run a Channels sample")
	flag.BoolVar(&libfunEnabled, "libfun", false, "Run some 'libfun' sample")
	flag.Parse()
}

func main() {
	if cacheEnabled { cacheSample() }
	if libfunEnabled { libfunSample() }
	if channelsEnabled { channelSample() }
}

func cacheSample() {

	cache := lrucache.Default()

	key := "test123"
	account := &Account{123, "test"}

	cache.Add(key, account)
	fmt.Println("Account added to cache")

	fmt.Println("Account retrieved from cache")
	account2 := cache.Get(key).(*Account)

	if &account2.id != &account.id {
		fmt.Printf("&account.id == %x vs &account2.id == %x", &account.id, &account2.id)
	} else {
		fmt.Println("Addresses of account.id match")
	}

	fmt.Println("Sleeping for 10 seconds")
	time.Sleep(10 * time.Second)

	fmt.Println("Manually purging")
	cache.Purge()
}

func libfunSample() {
	libfun.TestPointer()
	libfun.Hello()

	isFun := libfun.HaveFun()
	fmt.Printf("I am main.  I am %t\n", isFun)

	i := libfun.DeferFun()
	fmt.Printf("Value of i is %d\n", i)

	account := &libfun.Account{1, "test", time.Now()}
	fmt.Println(account.String())
}

func channelSample() {
	channels.TimeoutExample()
	channels.FixedMessagePump()
}
