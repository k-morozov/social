package api

import "net/http"

type SocialAPI struct {
}

var _ ServerAPI = &SocialAPI{}

func NewSocailAPI() ServerAPI {
	return &SocialAPI{}
}

func (srv *SocialAPI) Register(rw http.ResponseWriter, req *http.Request) {

}
