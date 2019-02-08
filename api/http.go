package api

import (
	"fmt"
	"html"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
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
		err := hystrix.Do(path, func() error {
			fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
			ha.httpSvc.Proxy(path, r)
			return nil
		}, func(err error) error {
			fmt.Printf("Fallback[%s] => got error: %+v\n", path, err)
			return nil
		})
		if err != nil {
			fmt.Printf("Warn: error in calling %s: %+v\n", path, err)
		}
	})
}
