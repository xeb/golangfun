package datastructs

import "fmt"
import "time"

type Account struct {
	Id          int64
	Name        string
	DateCreated time.Time
}

func CustomTypeExample() {
	fmt.Printf("-------------------------\n")
	account := &Account{Id: 1234, Name: "Mark", DateCreated: time.Now()}
	account2 := &Account{Id: 150056, Name: "KLim", DateCreated: time.Now().Add(24 * 6 * time.Hour)}
	newAccount := &account
	newAccount2 := &account2

	accounts := []*Account{account, *newAccount, account2, *newAccount2}

	fmt.Printf("&account == %s\n\n", &account)
	fmt.Printf("newAccount == %s\n\n", newAccount)
	fmt.Printf("&newAccount == %s\n\n", &newAccount)

	fmt.Printf("&account2 == %s\n\n", &account2)
	fmt.Printf("newAccount2 == %s\n\n", newAccount2)
	fmt.Printf("&newAccount2 == %s\n\n", &newAccount2)

	for i, a := range accounts {
		fmt.Printf("account[%d] (via a) == %s\n\n", i, a)         // variable
		fmt.Printf("*(&account[%d]) (via a) == %s\n\n", i, *(&a)) // value-at
		fmt.Printf("&account[%d] (via a) == %s\n\n", i, &a)       // address-of
	}

	for i := 0; i < len(accounts); i++ {
		fmt.Printf("account[%d] == %s\n\n", i, accounts[i])         // variable
		fmt.Printf("*(&account[%d]) == %s\n\n", i, *(&accounts[i])) // value-at
		fmt.Printf("&account[%d] == %s\n\n", i, &accounts[i])       // address-of
	}

	fmt.Printf("-------------------------\n")
}
