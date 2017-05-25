package virtualserver

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createVirtualServerAPI *CreateVirtualServerAPI
var newVirtualServerName = "exampleVirtualServer"

func createSetup() {
	newHTTPVirtualServer := HTTP{URIPath: "/download/private/status/check"}
	newBasicVirtualServer := Basic{
		Enabled:            false,
		DefaultTrafficPool: "pool_test_rui",
		Port:               80,
		Protocol:           "http",
	}
	newVirtualServerProperties := Properties{Basic: newBasicVirtualServer}
	newVirtualServer := VirtualServer{Properties: newVirtualServerProperties}

	createVirtualServerAPI = NewCreate(newVirtualServerName, newVirtualServer)
	createVirtualServerAPI.SetResponseObject("/download/private/status/check")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createVirtualServerAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(
		t,
		"/api/tm/3.8/config/active/virtual_servers/"+newVirtualServerName,
		createVirtualServerAPI.Endpoint(),
	)
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := `{"properties":{"basic":{"enabled":false,"pool":"pool_test_rui","port":80,"protocol":"http"},}}`
	jsonBytes, err := json.Marshal(createVirtualServerAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	getResponse := createVirtualServerAPI.GetResponse()
	assert.Equal(t, getResponse, "/download/private/status/check")

}

//TODO
func TestCreateUnMarshalling(t *testing.T) {
}
