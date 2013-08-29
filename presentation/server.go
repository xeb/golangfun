package presentation

import (
	"fmt"
	"html/template"
	"net/http"
)

// ridiculously long path name!
// var templatePath = "/Users/xeb/Dropbox/Projects/GoCode/src/github.com/xeb/golangfun/presentation/index.html"
var templatePath = "./presentation/index.html"
var port int = 8093

var contentMux = http.NewServeMux()

type Presentation struct {
	Name string
}

func Start() {
	contentMux.Handle("/", http.FileServer(http.Dir("./presentation/")))

	fmt.Printf("Listing on localhost:%d\n", port)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error :%s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling request %s\n", r.URL)
	url := r.URL.String()
	switch url {
	case "/":
		handleRootReq(w)
	default:
		contentMux.ServeHTTP(w, r)
		// fmt.Fprintf(w, "No handler for %s", url)
	}
}

func handleRootReq(w http.ResponseWriter) {
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
