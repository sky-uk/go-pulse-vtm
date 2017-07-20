package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetSingleTrafficIPGroupAPI base object.
type GetSingleTrafficIPGroupAPI struct {
	*rest.BaseAPI
}

// NewGetSingle returns a new object of GetTrafficIPGroupAPI.
func NewGetSingle(tipg string) *GetSingleTrafficIPGroupAPI {
	this := new(GetSingleTrafficIPGroupAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/traffic_ip_groups/"+tipg, nil, new(TrafficIPGroup), new(api.VTMError))
	return this
}

// GetResponse returns ResponseObject of GetSingleTrafficIPGroupAPI.
func (ga GetSingleTrafficIPGroupAPI) GetResponse() *TrafficIPGroup {
	return ga.ResponseObject().(*TrafficIPGroup)
}
