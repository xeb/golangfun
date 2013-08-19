package lrucacheTest

import (
	"fmt"
	"testing"
	"time"
	"github.com/xeb/golangfun/lrucache"
)

type Account struct {
	id int
	name string
}

func TestAddItem(t *testing.T) {
	key := "test123"
	account := &Account{123, "test"}

	lrucache.Add(key, account)
	account2 := lrucache.Get(key).(*Account)

	if &account2.id != &account.id {
		fmt.Printf("&account.id == %x vs &account2.id == %x", &account.id, &account2.id)
		t.Fatal("Address of original account's ID and the lrucache's received account ID are not equal")
	}
	
	time.Sleep(10 * time.Second)

	lrucache.Purge()

	t.Fatal("Done!")
}