package core

type ServiceDef struct {
	Proto string `yaml:"proto"`
	Host  string `yaml:"host"`
	Port  int16  `yaml:"port"`
}

func NewServiceDef(proto string, host string, port int16) ServiceDef {
	return ServiceDef{proto, host, port}
}

type ServicePath struct {
	Path         string `yaml:"path"`
	Method       string `yaml:"method"`
	MaxRequests  int    `yaml:"maxRequests"`
	Concurrency  int    `yaml:"concurrency"`
	Timeout      int    `yaml:"timeout"`
	ErrorPercent int    `yaml:"errorPercent"`
}

func (sp *ServicePath) Hash() string {
	return sp.Method + "__" + sp.Path
}

func NewServicePath(path string, method string, maxRequests int, concurrency int, timeout int, errorPercent int) ServicePath {
	return ServicePath{path, method, maxRequests, concurrency, timeout, errorPercent}
}
