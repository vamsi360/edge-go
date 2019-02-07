package api

import (
	"fmt"
	"html"
	"net/http"

	"github.com/edge-go/service"
)

type HttpApi struct {
	httpSvc *service.HttpService
}

func NewHttpApi(httpSvc *service.HttpService) *HttpApi {
	return &HttpApi{httpSvc}
}

func (ha *HttpApi) Handle(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
		ha.httpSvc.Proxy(path, r)
	})
}
