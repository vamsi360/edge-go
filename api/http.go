package api

import (
	"fmt"
	"html"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/edge-go/core"
	"github.com/edge-go/service"
)

type HttpApi struct {
	httpSvc *service.HttpService
}

func NewHttpApi(httpSvc *service.HttpService) *HttpApi {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 150
	return &HttpApi{httpSvc}
}

func (ha *HttpApi) Handle(path string, servicePath core.ServicePath) {
	hystrix.ConfigureCommand(path, hystrix.CommandConfig{
		Timeout:               servicePath.Timeout,
		MaxConcurrentRequests: servicePath.MaxRequests,
		ErrorPercentThreshold: servicePath.ErrorPercent,
	})

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		err := hystrix.Do(path, func() error {
			fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
			ha.httpSvc.Proxy(path, r)
			return nil
		}, func(err error) error {
			fmt.Printf("Fallback[%s] => got error: %+v\n", path, err)
			return err
		})
		if err != nil {
			fmt.Printf("Warn: error in calling %s: %+v\n", path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
