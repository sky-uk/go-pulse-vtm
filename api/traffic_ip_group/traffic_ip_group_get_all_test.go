package trafficIpGroups

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllTrafficIPGroupsAPI *GetAllTrafficIPGroupsAPI

//var getMonitorName string = "Simple HTTP"

func setupGetAll() {
	getAllTrafficIPGroupsAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllTrafficIPGroupsAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups", getAllTrafficIPGroupsAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	jsonContent := []byte("{\"children\":[{\"name\":\"test-group-1\",\"href\":\"/api/tm/3.8/config/active/traffic_ip_groups/test-group-1\"},{\"name\":\"test-group-2\",\"href\":\"/api/tm/3.8/config/active/traffic_ip_groups/test-group-2\"}]}")
	jsonErr := json.Unmarshal(jsonContent, getAllTrafficIPGroupsAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	assert.Len(t, getAllTrafficIPGroupsAPI.GetResponse().Children, 2)
	assert.Equal(t, "test-group-1", getAllTrafficIPGroupsAPI.GetResponse().Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-1", getAllTrafficIPGroupsAPI.GetResponse().Children[0].HRef)
	assert.Equal(t, "test-group-2", getAllTrafficIPGroupsAPI.GetResponse().Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-2", getAllTrafficIPGroupsAPI.GetResponse().Children[1].HRef)
}

func TestGetAllTrafficIPGroupsAPIGetResponse(t *testing.T) {
	setupGetAll()
	assert.IsType(t, getAllTrafficIPGroupsAPI.GetResponse(), &TrafficIPGroupList{})
}
