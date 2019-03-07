package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edge-go/api"
	"github.com/edge-go/config"
	"github.com/edge-go/service"
	"github.com/rcrowley/go-metrics"
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

	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	log.Fatal(http.ListenAndServe(":6001", nil))
	fmt.Printf("done..\n")
}
