package libfun

import "fmt"

func HaveFun() (b bool) {
	return true
}

func (a *Account) String() (s string) {
	s = fmt.Sprintf("Id=%d,Name=%s", a.Id, a.Name)
	return s
}


// Constant
var officePlace []string = []string{"Boston", "New York"}

type Office int

const (
	Boston Office=iota /* note that Office is used to initiaze the const */
	NewYork
)

func (o Office) String() string {
	return "Google, " + officePlace[o]
}

func Hello() {
	fmt.Printf("Hello, %s\n", Boston)
}