package trafficIpGroups

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllTrafficManagersAPI *GetAllTrafficManagersAPI
var trafficManagerChildren TrafficManagerChildren

func setupGetAllTrafficManagers() {
	getAllTrafficManagersAPI = NewGetTrafficManagerList()
	tmChild := TrafficMangerChild{Name: "h1ist01-v00.paas.d50.ovp.bskyb.com", HREF: "/api/tm/3.8/config/active/traffic_managers/h1ist01-v00.paas.d50.ovp.bskyb.com"}
	trafficManagerChildren.Children = append(trafficManagerChildren.Children, tmChild)
	getAllTrafficManagersAPI.SetResponseObject(&trafficManagerChildren)
}

func TestGetTrafficManagersMethod(t *testing.T) {
	setupGetAllTrafficManagers()
	assert.Equal(t, http.MethodGet, getAllTrafficManagersAPI.Method())
}

func TestGetTrafficManagersEndpoint(t *testing.T) {
	setupGetAllTrafficManagers()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_managers", getAllTrafficManagersAPI.Endpoint())
}

func TestGetTrafficManagersUnmarshalling(t *testing.T) {
	setupGetAllTrafficManagers()
	jsonContent := []byte(`{"children":[{"name":"h1ist01-v00.paas.d50.ovp.bskyb.com","href":"/api/tm/3.8/config/active/traffic_managers/h1ist01-v00.paas.d50.ovp.bskyb.com"}]}`)
	jsonErr := json.Unmarshal(jsonContent, getAllTrafficManagersAPI.ResponseObject())
	assert.Nil(t, jsonErr)
}

func TestGetTrafficManagersGetResponse(t *testing.T) {
	setupGetAllTrafficManagers()
	response := getAllTrafficManagersAPI.GetResponse()
	assert.IsType(t, getAllTrafficManagersAPI.GetResponse(), TrafficManagerChildren{})
	assert.Equal(t, "h1ist01-v00.paas.d50.ovp.bskyb.com", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_managers/h1ist01-v00.paas.d50.ovp.bskyb.com", response.Children[0].HREF)
}
