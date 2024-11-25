package json_api

import (
	"net/http"
	"strings"
)

func (p *operation) PossibleErr(a *ResponseStatus) *operation {
	// TODO
	return p
}
func (o *operation) __Writer() http.ResponseWriter { return o.writer }
func (o *operation) __Request() *http.Request      { return o.request }
func (o *operation) path() string {
	base := o.endpoint.basePath
	if strings.HasSuffix(base, "/") {
		base += "/"
	}
	return base + o.subPath
}
