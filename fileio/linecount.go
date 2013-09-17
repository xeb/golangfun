package fileio

import (
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

	for _, file := range files {
		fn := file.Name()
		switch {
		case file.IsDir():
			i = i + GetLineCount(path.Join(p, fn))
		case strings.HasSuffix(fn, ext):
			i = i + 1
			fmt.Printf("File Path %s\n", fn)
		}
	}
	return
}

func PrintErr(e error) {
	fmt.Printf("! ERROR ! %s", e)
	os.Exit(1)
}
