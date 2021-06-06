package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TextHandler все о запросе
func TextHandler(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	fmt.Fprintln(w, "Getwd():", dir)

	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}

	dat, err := ioutil.ReadFile("txt.txt")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(dat))
}
