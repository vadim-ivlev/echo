package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// EchoHandler все о запросе
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	s := headers(r)
	s += request(r)
	s += body(r)
	fmt.Fprint(w, s)
}

func headers(r *http.Request) string {
	ss := make([]string, 0)
	for name, headers := range r.Header {
		for _, header := range headers {
			// s += fmt.Sprintf("%v: %v \n", name, header)
			ss = append(ss, fmt.Sprintf("%v: %v", name, header))
		}
	}
	sort.Strings(ss)
	s := "HTTP Headers: ------------------------------------------------\n"
	s += strings.Join(ss, "\n")
	return s + "\n\n"
}

func body(r *http.Request) string {
	s := "HTTP Body: ---------------------------------------------------\n"
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	s += string(bodyBytes)
	// put it back
	// c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	return s + "\n\n"
}

func request(r *http.Request) string {
	s := "HTTP Request: ------------------------------------------------\n"
	s += "Request.Host: " + r.Host + "\n"
	s += "Request.Method: " + r.Method + "\n"
	s += "Request.RemoteAddr: " + r.RemoteAddr + "\n"
	s += "Request.RequestURI: " + r.RequestURI + "\n"
	s += "Request.Referer(): " + r.Referer() + "\n"
	s += "Request.ContentLength: " + strconv.FormatInt(r.ContentLength, 10) + "\n"
	s += "Request.URL.Host: " + r.URL.Host + "\n"
	s += "Request.URL.Path: " + r.URL.Path + "\n"
	s += "Request.URL.Hostname(): " + r.URL.Hostname() + "\n"
	s += "Request.URL.Port(): " + r.URL.Port() + "\n"
	s += "Request.URL.RequestURI(): " + r.URL.RequestURI() + "\n"
	s += "Request.URL.User.String(): " + r.URL.User.String() + "\n"
	s += "Request.URL.Query().Encode(): " + r.URL.Query().Encode() + "\n"
	s += "Request.URL.RawQuery: " + r.URL.RawQuery + "\n"
	return s + "\n"
}
