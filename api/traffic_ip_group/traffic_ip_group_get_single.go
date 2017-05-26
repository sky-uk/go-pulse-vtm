package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// GetTrafficIPGroupAPI base object.
type GetTrafficIPGroupAPI struct {
	*api.BaseAPI
}

// NewGetSingle returns a new object of GetTrafficIPGroupAPI.
func NewGetSingle(tipg string) *GetTrafficIPGroupAPI {
	this := new(GetTrafficIPGroupAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/traffic_ip_groups/"+tipg, nil, new(TrafficIPGroup))
	return this
}

// GetResponse returns ResponseObject of GetTrafficIPGroupAPI.
func (ga GetTrafficIPGroupAPI) GetResponse() *TrafficIPGroup {
	return ga.ResponseObject().(*TrafficIPGroup)
}
