package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// DeleteTrafficIPGroupAPI base object.
type DeleteTrafficIPGroupAPI struct {
	*rest.BaseAPI
}

// NewDelete returns a new object of DeleteTrafficIPGroupAPI.
func NewDelete(tipg string) *DeleteTrafficIPGroupAPI {
	this := new(DeleteTrafficIPGroupAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/traffic_ip_groups/"+tipg, nil, nil, new(api.VTMError))
	return this
}
