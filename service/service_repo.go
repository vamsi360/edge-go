package service

import (
	"fmt"

	"github.com/edge-go/core"
)

type ServiceRepo struct {
	defsMap map[core.ServiceDef]bool
	repoMap map[string]core.ServiceDef
}

func NewServiceRepo() *ServiceRepo {
	defsMap := make(map[core.ServiceDef]bool)
	repoMap := make(map[string]core.ServiceDef)
	return &ServiceRepo{defsMap, repoMap}
}

func (s *ServiceRepo) Register(serviceDef core.ServiceDef) {
	s.defsMap[serviceDef] = true
}

func (s *ServiceRepo) RegisterEdge(servicePath core.ServicePath, serviceDef core.ServiceDef) bool {
	if !s.defsMap[serviceDef] {
		fmt.Printf("Warning: ServiceDef %+v not found\n", serviceDef)
		return false
	}
	s.repoMap[servicePath.Hash()] = serviceDef
	return true
}

func (s *ServiceRepo) GetDef(servicePath core.ServicePath) core.ServiceDef {
	return s.repoMap[servicePath.Hash()]
}

func (s *ServiceRepo) GetDefs() []core.ServiceDef {
	defs := []core.ServiceDef{}
	for key, _ := range s.defsMap {
		defs = append(defs, key)
	}
	return defs
}
