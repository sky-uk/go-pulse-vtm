package sslServerKey

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// DeleteSSLServerKeyAPI : object used to call delete on monitor
type DeleteSSLServerKeyAPI struct {
	*rest.BaseAPI
}

// NewDelete : returns a new DeleteSSLServerKeyAPI object
func NewDelete(name string) *DeleteSSLServerKeyAPI {
	this := new(DeleteSSLServerKeyAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/ssl/server_keys/"+name, nil, nil, new(api.VTMError))
	return this
}
