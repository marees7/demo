package handler

import (
	"fmt"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go!")
}
