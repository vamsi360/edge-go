package api

import (
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/edge-go/core"
	"github.com/edge-go/service"
	"github.com/rcrowley/go-metrics"
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
		metrics.GetOrRegisterTimer("proxy.latency", nil).Time(func() {
			err := hystrix.Do(path, func() error {
				fmt.Fprintf(w, "URL: %q", html.EscapeString(r.URL.Path))
				err, resp := ha.httpSvc.Proxy(path, r)
				if err != nil {
					return err
				}
				if resp == nil {
					return errors.New("empty response from server")
				}
				bodyBytes, rErr := ioutil.ReadAll(resp.Body)
				if rErr != nil {
					return rErr
				}
				//w.WriteHeader(resp.StatusCode)
				w.Write(bodyBytes)
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
	})
}
