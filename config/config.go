package config

import (
	"io/ioutil"
	"log"

	"github.com/edge-go/core"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	edges []EdgeConf `yaml:"edges"`
}

type EdgeConf struct {
	tstore   ServiceConf `yaml:"tstore"`
	payments ServiceConf `yaml:"payments"`
}

type ServiceConf struct {
	serviceDef  core.ServiceDef  `yaml:"serviceDef"`
	servicePath core.ServicePath `yaml:"servicePath"`
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
