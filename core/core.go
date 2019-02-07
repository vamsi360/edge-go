package core

type ServiceDef struct {
	Proto string
	Host  string
	Port  int16
}

func NewServiceDef(proto string, host string, port int16) ServiceDef {
	return ServiceDef{proto, host, port}
}

type ServicePath struct {
	Path        string
	Method      string
	Headers     map[string]string
	Concurrency int16
	Timeout     int16
}

func (sp *ServicePath) Hash() string {
	return sp.Method + "__" + sp.Path
}

func NewServicePath(path string, method string, headers map[string]string, concurrency int16, timeout int16) ServicePath {
	return ServicePath{path, method, headers, concurrency, timeout}
}
