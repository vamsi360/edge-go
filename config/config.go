package config

import (
	"io/ioutil"
	"log"

	"github.com/edge-go/core"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Edges map[string]ServiceConf `yaml:"edges"`
}

type ServiceConf struct {
	EdgePath    string           `yaml:"edgePath"`
	ServiceDef  core.ServiceDef  `yaml:"serviceDef"`
	ServicePath core.ServicePath `yaml:"servicePath"`
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
