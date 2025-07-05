package api

import (
	"net/http"

	"github.com/rs/zerolog"
)

type SocialAPI struct {
	logger *zerolog.Logger
}

var _ ServerAPI = &SocialAPI{}

func NewSocailAPI(logger *zerolog.Logger) ServerAPI {
	return &SocialAPI{
		logger: logger,
	}
}

func (srv *SocialAPI) Register(rw http.ResponseWriter, req *http.Request) {

}
