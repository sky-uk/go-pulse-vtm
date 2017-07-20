package monitor

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// CreateMonitorAPI : create monitor API object
type CreateMonitorAPI struct {
	*rest.BaseAPI
}

// NewCreate : Create new monitor
func NewCreate(monitorName string, monitor Monitor) *CreateMonitorAPI {
	this := new(CreateMonitorAPI)
	requestPayLoad := new(Monitor)
	requestPayLoad.Properties.Basic = monitor.Properties.Basic
	requestPayLoad.Properties.HTTP = monitor.Properties.HTTP
	this.BaseAPI = rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/monitors/"+monitorName, requestPayLoad, new(Monitor), new(api.VTMError))
	return this
}

// GetResponse : get response object from created monitor
func (cma *CreateMonitorAPI) GetResponse() Monitor {
	return *cma.ResponseObject().(*Monitor)
}
