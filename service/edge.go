package service

import (
	"fmt"
	"net/http"

	"github.com/edge-go/core"
	"github.com/edge-go/util"
	"github.com/rcrowley/go-metrics"
)

type Edge struct {
}

func NewEdge() *Edge {
	return &Edge{}
}

func (e *Edge) Proxy(serviceDef core.ServiceDef, servicePath core.ServicePath, request *http.Request) (error, *http.Response) {
	contentType := request.Header.Get("Content-Type")
	url := util.MakeHttpUrl(serviceDef, servicePath)
	fmt.Printf("Making request to url: %s\n", url)

	var resp *http.Response
	var err error

	switch method := servicePath.Method; method {
	case "POST":
		metrics.GetOrRegisterTimer("post.latency", nil).Time(func() {
			resp, err = http.Post(url, contentType, request.Body)
		})
	case "GET":
		metrics.GetOrRegisterTimer("get.latency", nil).Time(func() {
			resp, err = http.Get(url)
		})
	}

	if err != nil {
		fmt.Printf("Error: in calling %s => %+v", url, err)
	} else if resp == nil {
		fmt.Printf("Error: Resp is nil")
	} else {
		fmt.Printf("Made a call to %s and got response: %+v\n", servicePath.Path, resp.StatusCode)
	}

	return err, resp
}
