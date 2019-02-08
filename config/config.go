package config

import (
	"io/ioutil"
	"log"

	"github.com/edge-go/core"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	serviceConfs []ServiceRootConf
}

type ServiceRootConf struct {
	serviceName string
	serviceConf ServiceConf
}

type ServiceConf struct {
	serviceDef  core.ServiceDef
	servicePath core.ServicePath
}

func ReadConf(fileName string) *Conf {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := &Conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
