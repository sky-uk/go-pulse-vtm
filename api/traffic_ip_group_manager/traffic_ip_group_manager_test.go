package trafficIpGroupManager

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testTrafficIPGroupManagerEndpoint = "/api/tm/3.8/config/active/traffic_managers/"

var getAllTrafficManagersAPI *rest.BaseAPI
var testTrafficManagerGetAllJSON []byte

func setupTrafficIPGroupManagersTest() {
	getAllTrafficManagersAPI = NewGetAll()
	testTrafficManagerGetAllJSON = []byte(`{"children":[{"name":"h1ist01-v00.paas.d50.ovp.bskyb.com","href":"/api/tm/3.8/config/active/traffic_managers/h1ist01-v00.paas.d50.ovp.bskyb.com"}]}`)
}

func TestGetTrafficManagersMethod(t *testing.T) {
	setupTrafficIPGroupManagersTest()
	assert.Equal(t, http.MethodGet, getAllTrafficManagersAPI.Method())
}

func TestGetTrafficManagersEndpoint(t *testing.T) {
	setupTrafficIPGroupManagersTest()
	assert.Equal(t, testTrafficIPGroupManagerEndpoint, getAllTrafficManagersAPI.Endpoint())
}

func TestGetTrafficManagersUnmarshalling(t *testing.T) {
	setupTrafficIPGroupManagersTest()
	jsonErr := json.Unmarshal(testTrafficManagerGetAllJSON, getAllTrafficManagersAPI.ResponseObject())
	response := getAllTrafficManagersAPI.ResponseObject().(*TrafficManagerChildren)

	assert.Nil(t, jsonErr)
	assert.Equal(t, 1, len(response.Children))
	assert.Equal(t, "h1ist01-v00.paas.d50.ovp.bskyb.com", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_managers/h1ist01-v00.paas.d50.ovp.bskyb.com", response.Children[0].HREF)
}
