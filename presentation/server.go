package presentation

import (
	"fmt"
	"html/template"
	"net/http"
)

var templatePath = "/Users/xeb/Dropbox/Projects/GoCode/src/github.com/xeb/golangfun/presentation/index.html"
var port int = 8093

type Presentation struct {
	Name string
}

func Start() {
	fmt.Printf("Listing on localhost:%d\n", port)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error :%s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Handling request %s\n", r.URL)
	t := template.New("index.html")
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		fmt.Fprintf(w, "ERROR %s", err)
		return
	}

	p := Presentation{Name: "TESTNAME"}
	err = t.Execute(w, p)
	if err != nil {
		fmt.Fprintf(w, "ERROR %s", err)
		return
	}

}
