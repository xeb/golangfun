package presentation

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	case "/compile":
		handleCompile(w, r)
	default:
		contentMux.ServeHTTP(w, r)
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

// NOTE: I am too lazy to write my own golang eval;
// or maybe I just haven't seen a great example online & don't want to figure it out
// either way, we are going to redirect the request to tour.golang.org/compile
func handleCompile(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	body, _ := r.Body.(io.ReadCloser)
	req := newProxyRequest(r.Header, *r.URL, r.Method, body)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	for key, _ := range resp.Header {
		w.Header().Add(key, resp.Header.Get(key))
	}

	w.WriteHeader(resp.StatusCode)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	writeBody(w, bodyBytes)

	fmt.Printf("Proxy received %d %s %s\n", resp.StatusCode, r.Method, r.URL.Path)
}

func newProxyRequest(origHeader http.Header, url url.URL, method string, body io.ReadCloser) (req *http.Request) {
	urlStr := "http://tour.golang.org/" + url.String()
	u, _ := url.Parse(urlStr)
	var header = make(http.Header)
	for key, _ := range origHeader {
		header.Add(key, origHeader.Get(key))
	}
	req = &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
		Host:       u.Host,
		Body:       body,
	}
	return
}

func writeBody(w http.ResponseWriter, bodyBytes []byte) {
	body := string(bodyBytes)
	fmt.Fprint(w, body)
}
