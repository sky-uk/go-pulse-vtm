package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetAllTrafficManagersAPI : base object.
type GetAllTrafficManagersAPI struct {
	*rest.BaseAPI
}

// NewGetTrafficManagerList : Get a list of traffic managers
func NewGetTrafficManagerList() *GetAllTrafficManagersAPI {
	this := new(GetAllTrafficManagersAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/traffic_managers", nil, new(TrafficManagerChildren), new(api.VTMError))
	return this
}

// GetResponse : get the response for GetAllTrafficManagersAPI
func (getAllTrafficManagers GetAllTrafficManagersAPI) GetResponse() TrafficManagerChildren {
	return *getAllTrafficManagers.ResponseObject().(*TrafficManagerChildren)
}
