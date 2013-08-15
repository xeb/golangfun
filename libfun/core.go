package libfun

import "fmt"

func HaveFun() (b bool) {
	return true
}

func (a *Account) String() (s string) {
	s = fmt.Sprintf("Id=%d,Name=%s", a.Id, a.Name)
	return s
}