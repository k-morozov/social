package api

import "net/http"

type ServerAPI interface {
	Register(rw http.ResponseWriter, req *http.Request)
}
