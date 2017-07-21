package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewCreate : Create new SSLServerKey
func NewCreate(name string, key SSLServerKey) *rest.BaseAPI {
	sslServerKeyCreateAPI := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/ssl/server_keys/"+name, key, new(SSLServerKey), new(api.VTMError))
	return sslServerKeyCreateAPI
}
