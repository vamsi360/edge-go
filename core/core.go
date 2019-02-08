package core

type ServiceDef struct {
	Proto string `yaml:"proto"`
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
}

func NewServiceDef(proto string, host string, port int16) ServiceDef {
	return ServiceDef{proto, host, port}
}

type ServicePath struct {
	Path         string `yaml:"path"`
	Method       string `yaml:"method"`
	Concurrency  int    `yaml:"concurrency"`
	Timeout      int    `yaml:"timeout"`
	ErrorPercent int    `yaml:"errorPercent"`
}

func (sp *ServicePath) Hash() string {
	return sp.Method + "__" + sp.Path
}

func NewServicePath(path string, method string, headers map[string]string, concurrency int16, timeout int16) ServicePath {
	return ServicePath{path, method, headers, concurrency, timeout}
}
