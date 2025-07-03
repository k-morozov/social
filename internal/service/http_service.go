package service

import "net/http"

type ServiceHTTP struct {
	server *http.Server
	// engine
	// storage
	// log
}

func NewServiceHTTP() (*ServiceHTTP, error) {
	srv := &ServiceHTTP{}
	return srv, nil
}
