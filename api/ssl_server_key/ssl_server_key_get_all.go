package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetAllSSLServerKeys base object.
type GetAllSSLServerKeys struct {
	*rest.BaseAPI
}

// NewGetAll returns a new object of GetAllSSLServerKeys.
func NewGetAll() *GetAllSSLServerKeys {
	this := new(GetAllSSLServerKeys)
	this.BaseAPI = rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/ssl/server_keys", nil, new(SSLServerKeysList), new(api.VTMError))
	return this
}

// GetResponse returns ResponseObject of SSLServerKeysList.
func (gam GetAllSSLServerKeys) GetResponse() SSLServerKeysList {
	return *gam.ResponseObject().(*SSLServerKeysList)
}
