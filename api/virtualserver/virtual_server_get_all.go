package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetAllVirtualServers : returns the list of all virtual servers
type GetAllVirtualServers struct {
	*rest.BaseAPI
}

// NewGetAll returns a new object of GetAllVirtualServers.
func NewGetAll() *GetAllVirtualServers {
	this := new(GetAllVirtualServers)
	this.BaseAPI = rest.NewBaseAPI(
		http.MethodGet,
		"/api/tm/3.8/config/active/virtual_servers",
		nil,
		new(VirtualServersList),
		new(api.VTMError),
	)
	return this
}

// GetResponse returns ResponseObject of GetAllVirtualServers.
func (gav GetAllVirtualServers) GetResponse() VirtualServersList {
	return *gav.ResponseObject().(*VirtualServersList)
}
