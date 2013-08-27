package main

import (
	"flag"
	"fmt"
	"github.com/xeb/golangfun/channels"
	"github.com/xeb/golangfun/datastructures"
	"github.com/xeb/golangfun/httpproxy"
	"github.com/xeb/golangfun/libfun"
	"github.com/xeb/golangfun/lrucache"
	"github.com/xeb/golangfun/presentation"
	"time"
)

type Account struct {
	id   int
	name string
}

var presentationEnabled bool
var linkedListEnabled bool
var httpProxyEnabled bool
var interfacesEnabled bool
var cacheEnabled bool
var channelsEnabled bool
var libfunEnabled bool

func init() {
	flag.BoolVar(&presentationEnabled, "presentation", false, "Runs the presentation")
	flag.BoolVar(&linkedListEnabled, "linkedlist", true, "Runs a linked list example")
	flag.BoolVar(&httpProxyEnabled, "httpproxy", false, "Runs an HTTP Proxy")
	flag.BoolVar(&interfacesEnabled, "interfaces", false, "Runs an example about interfaces")
	flag.BoolVar(&cacheEnabled, "cache", false, "Run a LRU Cache sample")
	flag.BoolVar(&channelsEnabled, "channels", false, "Run a Channels sample")
	flag.BoolVar(&libfunEnabled, "libfun", false, "Run some 'libfun' sample")
	flag.Parse()
}

func main() {
	if presentationEnabled {
		presentationSample()
	}
	if linkedListEnabled {
		linkedListSample()
	}
	if interfacesEnabled {
		interfacesSample()
	}
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

func presentationSample() {
	fmt.Println("Starting up presentation!")
	presentation.Start()
}

func linkedListSample() {
	ll := datastructs.NewLinkedList()
	fmt.Printf("%d elements initially\n", len(ll.Elements))

	times := 15
	for i := 0; i < times; i++ {
		acc := &Account{i, fmt.Sprintf("account%d", i)}
		ll.Append(acc)
		fmt.Printf("Appended account %d\n", i)
	}

	for i := 0; i < times; i++ {
		accE := ll.Get(i)
		if accE == nil {
			fmt.Printf("AccE is NIL; skipping\n")
			continue
		}
		fmt.Printf("Processing element %d\n", i)
		if accE.Prev != nil {
			fmt.Printf("\tPrev account is %d\n", accE.Prev.Item.(*Account).id)
		}
		if accE.Next != nil {
			fmt.Printf("\tNext account is %d\n", accE.Next.Item.(*Account).id)
		}
	}
}

func interfacesSample() {
	car := libfun.NewCar()
	fmt.Printf("Car==%s\n", car)
	car.PeddleToTheMetal()
	fmt.Println()

	truck := libfun.NewTruck()
	fmt.Printf("Truck==%s\n", truck)
	truck.PeddleToTheMetal()
	fmt.Println()

	fmt.Printf("Calling PunchIt on Car...\n")
	libfun.PunchIt(car)

	fmt.Printf("Calling PunchIt on Truck...\n")
	libfun.PunchIt(truck)
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
