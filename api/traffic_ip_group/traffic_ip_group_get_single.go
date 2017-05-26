package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// GetTrafficIpGroupAPI base object.
type GetTrafficIpGroupAPI struct {
	*api.BaseAPI
}

// NewGetSingle returns a new object of GetTrafficIpGroupAPI.
func NewGetSingle(tipg string) *GetTrafficIpGroupAPI {
	this := new(GetTrafficIpGroupAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/traffic_ip_groups/"+tipg, nil, new(TrafficIPGroup))
	return this
}

func (ga GetTrafficIpGroupAPI) GetResponse() *TrafficIPGroup {
	return ga.ResponseObject().(*TrafficIPGroup)
}
