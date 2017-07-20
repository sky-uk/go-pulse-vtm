package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// DeleteVirtualServerAPI : object used to call delete on monitor
type DeleteVirtualServerAPI struct {
	*rest.BaseAPI
}

// NewDelete : returns a new DeleteVirtualServerAPI object
func NewDelete(name string) *DeleteVirtualServerAPI {
	this := new(DeleteVirtualServerAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/virtual_servers/"+name, nil, nil, new(api.VTMError))
	return this
}
