package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// CreateVirtualServerAPI : Create VirtualServer API
type CreateVirtualServerAPI struct {
	*api.BaseAPI
}

// NewCreate : Create new virtualServer
func NewCreate(virtualServerName string, virtualServer VirtualServer) *CreateVirtualServerAPI {
	this := new(CreateVirtualServerAPI)
	requestPayLoad := new(VirtualServer)
	requestPayLoad.Properties.Basic = virtualServer.Properties.Basic
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/virtualServers/"+virtualServerName, requestPayLoad, new(string))
	return this
}

// GetResponse : get response object from created virtualServer
func (cma CreateVirtualServerAPI) GetResponse() string {
	return cma.ResponseObject().(string)
}
