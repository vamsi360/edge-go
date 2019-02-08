package service

import "net/http"

type HttpService struct {
	repo *ServiceRepo
	edge *Edge
}

func NewHttpService(repo *ServiceRepo, edge *Edge) *HttpService {
	return &HttpService{repo, edge}
}

func (hs *HttpService) Proxy(edgePath string, request *http.Request) (error, *http.Response) {
	sPath := hs.repo.GetPath(edgePath)
	sDef := hs.repo.GetDef(sPath)
	return hs.edge.Proxy(sDef, sPath, request)
}
