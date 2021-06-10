package echo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func Http(r *http.Request) string {
	s := "HTTP >>>>\n"
	s += "\nRequest: ------------------------------------------------\n"
	s += Request(r)
	s += "\nHeaders: ------------------------------------------------\n"
	s += Headers(r)
	s += "\nBody: ---------------------------------------------------\n"
	s += Body(r)
	s += "HTTP <<<<\n"
	return s
}

func Headers(r *http.Request) string {
	ss := make([]string, 0)
	for name, headers := range r.Header {
		for _, header := range headers {
			// s += fmt.Sprintf("%v: %v \n", name, header)
			ss = append(ss, fmt.Sprintf("%v: %v", name, header))
		}
	}
	sort.Strings(ss)
	s := strings.Join(ss, "\n")
	return s + "\n"
}

func Body(r *http.Request) string {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	s := string(bodyBytes)
	// put it back
	r.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	return s + "\n"
}

func Request(r *http.Request) string {
	s := "Request.Host: " + r.Host + "\n"
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
	s += "Request.URL.Query().Encode(): " + r.URL.RawQuery + "\n"
	return s
}
