package api

import (
	"fmt"
	"github.com/vadim-ivlev/echo/pkg/echo"
	"net/http"
)

// EchoHttp prints out everithing about request
func EchoHttp(w http.ResponseWriter, r *http.Request) {
	s := echo.Http(r)
	fmt.Fprint(w, s)
}
