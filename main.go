package main

import (
	"flag"
	"fmt"
	"github.com/xeb/golangfun/channels"
	"github.com/xeb/golangfun/httpproxy"
	"github.com/xeb/golangfun/libfun"
	"github.com/xeb/golangfun/lrucache"
	"time"
)

type Account struct {
	id   int
	name string
}

var httpProxyEnabled bool
var cacheEnabled bool
var channelsEnabled bool
var libfunEnabled bool

func init() {
	flag.BoolVar(&httpProxyEnabled, "httpproxy", true, "Runs an HTTP Proxy")
	flag.BoolVar(&cacheEnabled, "cache", false, "Run a LRU Cache sample")
	flag.BoolVar(&channelsEnabled, "channels", false, "Run a Channels sample")
	flag.BoolVar(&libfunEnabled, "libfun", false, "Run some 'libfun' sample")
	flag.Parse()
}

func main() {
	if httpProxyEnabled {
		httpProxySample()
	}
	if cacheEnabled {
		cacheSample()
	}
	if libfunEnabled {
		libfunSample()
	}
	if channelsEnabled {
		channelSample()
	}
}

func httpProxySample() {
	fmt.Println("Starting httpproxy server...")
	httpproxy.Start()
}

func cacheSample() {

	cache := lrucache.New(10)

	var account0 *Account
	for i := 0; i < 12; i++ {
		key := fmt.Sprintf("test%d", i)
		account := &Account{i, key}
		if i == 0 {
			account0 = account
		}
		cache.Add(key, account)
		fmt.Printf("Account %s added to cache\n", key)

		// Keep getting account0 and account1 to keep in the cache
		cache.Get("test0")
		cache.Get("test1")
	}

	account0b := cache.Get("test0").(*Account)
	fmt.Println("Account0 retrieved from cache")

	if &account0b.id != &account0.id {
		fmt.Printf("&account0.id == %x vs &account0b.id == %x\n", &account0.id, &account0b.id)
	} else {
		fmt.Println("Addresses of account0.id match!!")
	}

	fmt.Println("Manually evicting")
	cache.Evict()
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
