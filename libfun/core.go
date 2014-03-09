package libfun

import "fmt"

func HaveFun() (b bool) {
	return true
}

func (a *Account) String() (s string) {
	s = fmt.Sprintf("Id=%d,Name=%s", a.Id, a.Name)
	return s
}

// var sampleMap map[string][]string

// Constant
var officePlace []string = []string{"Boston", "New York"}

type Office int

const (
	Boston Office = iota /* note that Office is used to initiaze the const */
	NewYork
)

func (o Office) String() string {
	return "Google, " + officePlace[o]
}

func Hello() {
	fmt.Printf("Hello, %s\n", Boston)
}

func TestPointer() {
	var c *int = getPtr()
	fmt.Println("------")
	fmt.Printf("c Address is 0x%x\n", c)
	fmt.Printf("c Value is %d\n", *c)
	fmt.Printf("Address of c value is 0x%x\n", &c)
	fmt.Println("------")

	*c = 125
	fmt.Println("------")
	fmt.Printf("c Address is 0x%x\n", c)
	fmt.Printf("c Value is %d\n", *c)
	fmt.Printf("Address of c value is 0x%x\n", &c)
	fmt.Println("------")

	var d int = 130
	c = &d
	fmt.Println("------")
	fmt.Printf("c Address is 0x%x\n", c)
	fmt.Printf("c Value is %d\n", *c)
	fmt.Printf("Address of c value is 0x%x\n", &c)
	fmt.Println("------")

}

func getPtr() *int {
	var a int = 123
	var b *int = &a
	return b
}
