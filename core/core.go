package core

type ServiceDef struct {
	proto string
	host  string
	port  int16
}

func NewServiceDef(proto string, host string, port int16) ServiceDef {
	return ServiceDef{proto, host, port}
}

type ServicePath struct {
	path        string
	concurrency int16
	timeout     int16
}
