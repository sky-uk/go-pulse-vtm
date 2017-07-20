package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// CreateSSLServerKeyAPI : Create Monitor API
type CreateSSLServerKeyAPI struct {
	*rest.BaseAPI
}

// NewCreate : Create new SSLServerKey
func NewCreate(name string, key SSLServerKey) *CreateSSLServerKeyAPI {
	this := new(CreateSSLServerKeyAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/ssl/server_keys/"+name, key, new(SSLServerKey), new(api.VTMError))
	return this
}

// GetResponse : get response object from create SSLServerKey api call.
func (cma CreateSSLServerKeyAPI) GetResponse() SSLServerKey {
	return *cma.ResponseObject().(*SSLServerKey)
}
