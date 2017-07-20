package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// UpdateVirtualServerAPI : object we use to update a virtual server
type UpdateVirtualServerAPI struct {
	*rest.BaseAPI
}

// NewUpdate : creates a new object of type UpdateVirtualServerAPI
func NewUpdate(name string, virtualServer VirtualServer) *UpdateVirtualServerAPI {
	this := new(UpdateVirtualServerAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/virtual_servers/"+name, virtualServer, new(VirtualServer), new(api.VTMError))
	return this
}

// GetResponse : returns the response object from UpdateVirtualServerAPI
func (updateVirtualServerAPI UpdateVirtualServerAPI) GetResponse() VirtualServer {
	return *updateVirtualServerAPI.ResponseObject().(*VirtualServer)
}
