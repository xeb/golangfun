package httpproxy

import (
	"fmt"
	"net/http"
	"html"
)

func Start() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe("0.0.0.0:8092", nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error :%s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}