package service

import "github.com/edge-go/core"

type ServiceRepo struct {
	repoMap map[string]core.ServiceDef
}

func NewServiceRepo() *ServiceRepo {
	repoMap := make(map[string]core.ServiceDef)
	return &ServiceRepo{repoMap}
}

func (s *ServiceRepo) Register(edgePath string, serviceDef core.ServiceDef) {
	s.repoMap[edgePath] = serviceDef
}

func (s *ServiceRepo) Get(edgePath string) core.ServiceDef {
	return s.repoMap[edgePath]
}

func (s *ServiceRepo) GetAll() []core.ServiceDef {
	defs := []core.ServiceDef{}
	for _, value := range s.repoMap {
		defs = append(defs, value)
	}
	return defs
}
