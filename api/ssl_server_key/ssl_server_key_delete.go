package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewDelete : returns a new DeleteSSLServerKeyAPI object
func NewDelete(name string) *rest.BaseAPI {
	sslServerKeyDeleteAPI := rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/ssl/server_keys/"+name, nil, nil, new(api.VTMError))
	return sslServerKeyDeleteAPI
}
