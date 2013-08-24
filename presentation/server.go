package presentation

import (
	"fmt"
	"net/http"
	"time"
	// "net/url"
	// "io/ioutil"
	// "strconv"
	// "strings"
	// "io"
)

var port int = 8093

func Start() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error :%s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! %s", time.Now())
}
