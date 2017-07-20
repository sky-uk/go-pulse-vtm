package monitor

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetAllMonitors base object.
type GetAllMonitors struct {
	*rest.BaseAPI
}

// NewGetAll returns a new object of GetAllMonitors.
func NewGetAll() *GetAllMonitors {
	this := new(GetAllMonitors)
	this.BaseAPI = rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/monitors", nil, new(MonitorsList), new(api.VTMError))
	return this
}

// GetResponse returns ResponseObject of GetAllMonitors.
func (gam GetAllMonitors) GetResponse() MonitorsList {
	return *gam.ResponseObject().(*MonitorsList)
}
