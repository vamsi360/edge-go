package service

import (
	"fmt"

	"github.com/edge-go/core"
)

type ServiceRepo struct {
	defsMap  map[core.ServiceDef]bool
	pathsMap map[string]core.ServicePath //edgePath to ServicePath
	repoMap  map[string]core.ServiceDef
}

func NewServiceRepo() *ServiceRepo {
	defsMap := make(map[core.ServiceDef]bool)
	pathsMap := make(map[string]core.ServicePath)
	repoMap := make(map[string]core.ServiceDef)
	return &ServiceRepo{defsMap, pathsMap, repoMap}
}

func (s *ServiceRepo) Register(serviceDef core.ServiceDef) {
	s.defsMap[serviceDef] = true
}

func (s *ServiceRepo) RegisterEdge(edgePath string, servicePath core.ServicePath, serviceDef core.ServiceDef) bool {
	if !s.defsMap[serviceDef] {
		fmt.Printf("Warning: ServiceDef %+v not found\n", serviceDef)
		return false
	}
	s.pathsMap[edgePath] = servicePath
	s.repoMap[servicePath.Hash()] = serviceDef
	return true
}

func (s *ServiceRepo) GetDef(servicePath core.ServicePath) core.ServiceDef {
	return s.repoMap[servicePath.Hash()]
}

func (s *ServiceRepo) GetPath(edgePath string) core.ServicePath {
	return s.pathsMap[edgePath]
}

func (s *ServiceRepo) GetPaths() []core.ServicePath {
	paths := []core.ServicePath{}
	for _, path := range s.pathsMap {
		paths = append(paths, path)
	}
	return paths
}

func (s *ServiceRepo) GetDefs() []core.ServiceDef {
	defs := []core.ServiceDef{}
	for key, _ := range s.defsMap {
		defs = append(defs, key)
	}
	return defs
}
