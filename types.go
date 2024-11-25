package json_api

import (
	"net/http"
	"os"
)

type endpoint struct {
	basePath string
	docsFile *os.File
}

type None struct{}
type handleCallbackP[P any, R Response] func(P) R
type handleCallbackPQ[P, Q any, R Response] func(P, Q) R
type handleCallbackPQB[P, Q, B any, R Response] func(P, Q, B) R
type handleCallback[P, Q, B, H any, R Response] func(P, Q, B, H) R

type operation struct {
	*endpoint
	request     *http.Request
	writer      http.ResponseWriter
	subPath     string
	method      string
	description string
}
