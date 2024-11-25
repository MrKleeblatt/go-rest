package json_api

import (
	"net/http"
	"os"

	"github.com/MrKleeblatt/go-log"
)

var logger *log.Logger = log.New(os.Stdout).WithDebug()

func Handle[P, Q, B, H any, R Response](o *operation, cb handleCallback[P, Q, B, H, R]) error {
	path := o.path()
	logger.Debug("creating", o.method, "on", path)
	// TODO: documentation generation
	http.Handle(o.path(), &handleWrapper[P, Q, B, H, R]{operation: o, handleCallback: cb})
	return nil
}

func HandleP[P any, R Response](o *operation, cb handleCallbackP[P, R]) error {
	return Handle(o, func(p P, _ None, _ None, _ None) R { return cb(p) })
}

func HandlePQ[P, Q any, R Response](o *operation, cb handleCallbackPQ[P, Q, R]) error {
	return Handle(o, func(p P, q Q, _ None, _ None) R { return cb(p, q) })
}

func HandlePQB[P, Q, B any, R Response](o *operation, cb handleCallbackPQB[P, Q, B, R]) error {
	return Handle(o, func(p P, q Q, b B, _ None) R { return cb(p, q, b) })
}

func Endpoint(basePath, description string) *endpoint { return &endpoint{basePath: basePath} }

func (p *endpoint) Get(path, description string) *operation {
	return p.register("GET", path, description)
}
func (p *endpoint) Head(path, description string) *operation {
	return p.register("HEAD", path, description)
}
func (p *endpoint) Post(path, description string) *operation {
	return p.register("POST", path, description)
}
func (p *endpoint) Put(path, description string) *operation {
	return p.register("PUT", path, description)
}
func (p *endpoint) Delete(path, description string) *operation {
	return p.register("DELETE", path, description)
}
func (p *endpoint) Connect(path, description string) *operation {
	return p.register("CONNECT", path, description)
}
func (p *endpoint) Options(path, description string) *operation {
	return p.register("OPTIONS", path, description)
}
func (p *endpoint) Trace(path, description string) *operation {
	return p.register("TRACE", path, description)
}
func (p *endpoint) Patch(path, description string) *operation {
	return p.register("PATCH", path, description)
}
