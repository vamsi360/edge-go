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
	Concurrency  int16  `yaml:"concurrency"`
	Timeout      int16  `yaml:"timeout"`
	ErrorPercent int16  `yaml:"errorPercent"`
}

func (sp *ServicePath) Hash() string {
	return sp.Method + "__" + sp.Path
}

func NewServicePath(path string, method string, concurrency int16, timeout int16, errorPercent int16) ServicePath {
	return ServicePath{path, method, concurrency, timeout, errorPercent}
}
