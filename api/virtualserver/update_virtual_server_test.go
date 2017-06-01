package virtualserver

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateVirtualServerAPI *UpdateVirtualServerAPI
var updateVirtualServerName = "updateVirtualServer"

func setup() {
	newBasicVirtualServer := Basic{
		Enabled:  false,
		Pool:     "pool_test_rui",
		Port:     80,
		Protocol: "http",
	}
	newVirtualServerProperties := Properties{Basic: newBasicVirtualServer}
	newVirtualServer := VirtualServer{Properties: newVirtualServerProperties}

	updateVirtualServerAPI = NewUpdate(updateVirtualServerName, newVirtualServer)
	updateVirtualServerAPI.SetResponseObject("/private/status/check")
}

func TestUpdateMethod(t *testing.T) {
	setup()
	assert.Equal(t, http.MethodPut, updateVirtualServerAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setup()
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/"+updateVirtualServerName, updateVirtualServerAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	setup()
	expectedJSON := `{"properties":{"basic":{"enabled":false,"pool":"pool_test_rui","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{"ssl_data":false},"gzip":{},"ssl":{}}}`
	jsonBytes, err := json.Marshal(updateVirtualServerAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestUpdateGetResponse(t *testing.T) {
	setup()
	getResponse := updateVirtualServerAPI.GetResponse()
	assert.Equal(t, getResponse, "/private/status/check")

}
