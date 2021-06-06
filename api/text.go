package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// TextHandler все о запросе
func TextHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("api/text.txt")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(dat))
}
