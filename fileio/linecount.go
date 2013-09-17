package fileio

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var ext string = ".go"

func GetLineCount(p string) (i int) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		PrintErr(err)
	}

	// var c int
	for _, file := range files {
		fn := file.Name()
		fp := path.Join(p, fn)
		switch {
		case file.IsDir():
			i = i + GetLineCount(fp)
		case strings.HasSuffix(fn, ext):
			c := ReadLineCount(fp)
			i = i + c
			fmt.Printf("File Path %s has %d\n", fn, c)
		}
	}
	return
}

func ReadLineCount(f string) (i int) {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// t := scanner.Text()
		// fmt.Printf("\t\tText==%s\n", t)
		i = i + 1
	}
	return i
}

func PrintErr(e error) {
	fmt.Printf("! ERROR ! %s", e)
	os.Exit(1)
}
