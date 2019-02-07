package service

import (
	"fmt"
	"net/http"

	"github.com/edge-go/core"
	"github.com/edge-go/util"
)

type Edge struct {
}

func NewEdge() *Edge {
	return &Edge{}
}

func (e *Edge) Proxy(serviceDef core.ServiceDef, servicePath core.ServicePath, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	url := util.MakeHttpUrl(serviceDef, servicePath)
	fmt.Printf("Making request to url: %s\n", url)

	var resp *http.Response
	var err error

	switch method := servicePath.Method; method {
	case "POST":
		resp, err = http.Post(url, contentType, request.Body)
	case "GET":
		resp, err = http.Get(url)
	}

	if err != nil {
		fmt.Printf("Error: in calling %s\n", url)
	} else {
		fmt.Printf("Made a call to %s and got response: %d\n", servicePath.Path, resp.StatusCode)
	}
}
