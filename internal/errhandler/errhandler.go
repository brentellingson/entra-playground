package errhandler

import (
	"net/http"
)

// Abort aborts the request with an error.
func Abort(w http.ResponseWriter, code int, err error) {
	http.Error(w, err.Error(), code)
}
