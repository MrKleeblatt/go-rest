package json_api

import "os"


func (p *endpoint) register(method, subPath, description string) *operation {
	return &operation{endpoint: p, method: method, subPath: subPath, description: description}
}

func (p *endpoint) Documentation(f *os.File) *endpoint {
	p.docsFile = f
	return p
}
