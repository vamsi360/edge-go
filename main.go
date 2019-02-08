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

	// tstoreEP := "/changes"
	// tstoreSD := core.NewServiceDef("http", "localhost", 18080)
	// serviceRepo.Register(tstoreSD)

	// tstoreSP := core.NewServicePath("v1/entity/changes/meta_13", "GET", 1, 1000, 5)
	// serviceRepo.RegisterEdge(tstoreEP, tstoreSP, tstoreSD)

	// sdef1 := serviceRepo.GetDef(tstoreSP)
	// fmt.Printf("Repo1: %v\n", sdef1)

	// sdefs := serviceRepo.GetDefs()
	// fmt.Printf("Registered ServiceDefs: %v\n", sdefs)

	// sPaths := serviceRepo.GetPaths()
	// fmt.Printf("Registered ServicePaths: %v\n", sPaths)

	// tstoreReq, err := http.NewRequest(tstoreSP.Method, tstoreSP.Path, nil)
	// if err != nil {
	// 	fmt.Printf("Error: Unable to create request %+v\n", tstoreReq)
	// }

	edge := service.NewEdge()
	httpSvc := service.NewHttpService(serviceRepo, edge)
	httpApi := api.NewHttpApi(httpSvc)

	conf := config.ReadConf("config.yaml")
	//fmt.Printf("Config: %+v\n", conf)

	for key, svcConf := range conf.Edges {
		fmt.Printf("Key: %+v and Val: %+v\n", key, svcConf)
		serviceRepo.Register(svcConf.ServiceDef)
		serviceRepo.RegisterEdge(svcConf.EdgePath, svcConf.ServicePath, svcConf.ServiceDef)
		httpApi.Handle(svcConf.EdgePath)
	}

	log.Fatal(http.ListenAndServe(":6001", nil))
	fmt.Printf("done..\n")
}
