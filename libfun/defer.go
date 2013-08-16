package libfun

func DeferFun() (i int) {
	defer func () { i = 2 }()
	return 1
}