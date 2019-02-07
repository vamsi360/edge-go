package main

import (
	"fmt"
	"net/http"

	"github.com/edge-go/core"
	"github.com/edge-go/service"
)

func main() {
	fmt.Printf("Starting app..\n")

	serviceRepo := service.NewServiceRepo()

	tstoreSD := core.NewServiceDef("http", "localhost", 18080)
	serviceRepo.Register(tstoreSD)

	tstoreSPHeaders := map[string]string{}
	tstoreSP := core.NewServicePath("v1/entity/changes/meta_13", "GET", tstoreSPHeaders, 1, 1000)
	serviceRepo.RegisterEdge("changes", tstoreSP, tstoreSD)

	sdef1 := serviceRepo.GetDef(tstoreSP)
	fmt.Printf("Repo1: %v\n", sdef1)

	sdefs := serviceRepo.GetDefs()
	fmt.Printf("Registered ServiceDefs: %v\n", sdefs)

	sPaths := serviceRepo.GetPaths()
	fmt.Printf("Registered ServicePaths: %v\n", sPaths)

	tstoreReq, err := http.NewRequest(tstoreSP.Method, tstoreSP.Path, nil)
	if err != nil {
		fmt.Printf("Error: Unable to create request %+v\n", tstoreReq)
	}

	edge := service.NewEdge()
	httpSvc := service.NewHttpService(serviceRepo, edge)
	httpSvc.Proxy("changes", tstoreReq)

	fmt.Printf("done..")
}
