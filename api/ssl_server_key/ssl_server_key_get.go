package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewGet : returns the SSLServerKey details
func NewGet(name string) *rest.BaseAPI {
	SSLServerKeyGetAPI := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/ssl/server_keys/"+name, nil, new(SSLServerKey), new(api.VTMError))
	return SSLServerKeyGetAPI
}
