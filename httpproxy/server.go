package httpproxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var port int = 8093
var portStr string = strconv.Itoa(port)
var baseUrl string = "http://www.google.com"
var replacements map[string][]string = map[string][]string{
	"http://127.0.0.1:" + portStr: []string{"http://www.google.com", "https://www.google.com", "http://google.com", "https://www.google.com"},
	"//127.0.0.1:" + portStr:      []string{"//www.google.com"},
}

func Start() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error :%s\n", err)
	}
}

func newProxyRequest(origHeader http.Header, url url.URL, method string, body io.ReadCloser) (req *http.Request) {
	urlStr := baseUrl + url.String()
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
	for key, _ := range replacements {
		for i := range replacements[key] {
			body = strings.Replace(body, replacements[key][i], key, 100)
		}
	}
	fmt.Fprint(w, body)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("Handling %d %s %s\n", resp.StatusCode, r.Method, r.URL.Path)
}
