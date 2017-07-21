package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewGetAll returns a new object of GetAllSSLServerKeys.
func NewGetAll() *rest.BaseAPI {
	SSLServerKeyGetAllAPI := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/ssl/server_keys", nil, new(SSLServerKeysList), new(api.VTMError))
	return SSLServerKeyGetAllAPI
}
