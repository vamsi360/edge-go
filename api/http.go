package api

import (
	"fmt"
	"html"
	"net/http"
)

type HttpApi struct {
}

func (ha *HttpApi) handle(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
	})
}
