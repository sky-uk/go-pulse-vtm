package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// CreateTrafficIPGroupAPI base object.
type CreateTrafficIPGroupAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateTrafficIPGroupA.
func NewCreate(name string, requestPayload TrafficIPGroup) *CreateTrafficIPGroupAPI{
	this := new(CreateTrafficIPGroupAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/traffic_ip_groups/"+name, requestPayload, new(TrafficIPGroup))
	return this
}

func (ga CreateTrafficIPGroupAPI) GetResponse() *TrafficIPGroup {
	return ga.ResponseObject().(*TrafficIPGroup)
}