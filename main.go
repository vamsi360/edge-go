package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"

	"github.com/edge-go/api"
	"github.com/edge-go/core"
	"github.com/edge-go/service"

	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Printf("Starting app..\n")

	serviceRepo := service.NewServiceRepo()

	tstoreEP := "/changes"
	tstoreSD := core.NewServiceDef("http", "localhost", 18080)
	serviceRepo.Register(tstoreSD)

	tstoreSPHeaders := map[string]string{}
	tstoreSP := core.NewServicePath("v1/entity/changes/meta_13", "GET", tstoreSPHeaders, 1, 1000)
	serviceRepo.RegisterEdge(tstoreEP, tstoreSP, tstoreSD)

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

	httpApi := api.NewHttpApi(httpSvc)
	httpApi.Handle(tstoreEP)

	//httpSvc.Proxy(tstoreEP, tstoreReq)

	//log.Fatal(http.ListenAndServe(":6001", nil))
	//fmt.Printf("done..\n")

	var c conf
    c.getConf()
    fmt.Println(c)
}

type conf struct {
	Services []struct {
		Tstore struct {
			ServiceDef struct {
				Proto string `yaml:"proto"`
				Host  string `yaml:"host"`
				Port  int    `yaml:"port"`
			} `yaml:"serviceDef"`
			ServicePath struct {
				Path         string `yaml:"path"`
				Method       string `yaml:"method"`
				Concurrency  int    `yaml:"concurrency"`
				Timeout      int    `yaml:"timeout"`
				ErrorPercent int    `yaml:"errorPercent"`
			} `yaml:"servicePath"`
		} `yaml:"tstore,omitempty"`
		Payments struct {
			ServiceDef struct {
				Proto string `yaml:"proto"`
				Host  string `yaml:"host"`
				Port  int    `yaml:"port"`
			} `yaml:"serviceDef"`
			ServicePath struct {
				Path         string `yaml:"path"`
				Method       string `yaml:"method"`
				Concurrency  int    `yaml:"concurrency"`
				Timeout      int    `yaml:"timeout"`
				ErrorPercent int    `yaml:"errorPercent"`
			} `yaml:"servicePath"`
		} `yaml:"payments,omitempty"`
	} `yaml:"services"`
}


func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}