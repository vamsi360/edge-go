package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edge-go/api"
	"github.com/edge-go/config"
	"github.com/edge-go/service"
)

func main() {
	fmt.Printf("Starting app..\n")

	serviceRepo := service.NewServiceRepo()
	edge := service.NewEdge()
	httpSvc := service.NewHttpService(serviceRepo, edge)
	httpApi := api.NewHttpApi(httpSvc)

	conf := config.ReadConf("config.yaml")
	fmt.Printf("Config: %+v\n", conf)

	for key, svcConf := range conf.Edges {
		fmt.Printf("Key: %+v and Val: %+v\n", key, svcConf)
		serviceRepo.Register(svcConf.ServiceDef)
		serviceRepo.RegisterEdge(svcConf.EdgePath, svcConf.ServicePath, svcConf.ServiceDef)
		httpApi.Handle(svcConf.EdgePath, svcConf.ServicePath)
	}

	log.Fatal(http.ListenAndServe(":6001", nil))
	fmt.Printf("done..\n")
}
