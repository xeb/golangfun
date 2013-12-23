package fileio

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var ext []string = []string{".php", ".html", ".js",".cs", ".go", ".java", ".cpp", ".h", ".xml", ".config"}

func GetLineCount(p string) (i int) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		PrintErr(err)
	}

	w := 0
	t := make(chan int)
	for _, file := range files {
		fn := file.Name()
		fp := path.Join(p, fn)

		if file.IsDir() {
			w = w + 1
			go func() {
				t <- GetLineCount(fp)
			}()
			continue
		}

		for _, item := range ext {
			if strings.HasSuffix(fn, item) {
				w = w + 1
				go func() {
					c := ReadLineCount(fp)
					// fmt.Printf("File Path %s has %d\n", fp, c)
					t <- c
				}()
			}
		}
	}

	for j := 0; j < w; j++ {
		i = i + <-t
	}
	close(t)
	return
}

func ReadLineCount(f string) (i int) {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i = i + 1
	}
	return i
}

func PrintErr(e error) {
	fmt.Printf("! ERROR ! %s", e)
	os.Exit(1)
}
