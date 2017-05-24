package monitor

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

type CreateMonitorAPI struct {
	*api.BaseAPI
}

func NewCreate(monitorName string, monitor Monitor) *CreateMonitorAPI {
	this := new(CreateMonitorAPI)
	requestPayLoad := new(Monitor)
	requestPayLoad.Properties.Basic = monitor.Properties.Basic
	requestPayLoad.Properties.Http = monitor.Properties.Http
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/monitors/"+monitorName, requestPayLoad, new(string))
	return this
}

func (cma CreateMonitorAPI) GetResponse() string {
	return cma.ResponseObject().(string)
}