package virtualserver

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateVirtualServerAPI *UpdateVirtualServerAPI
var updateVirtualServerName = "updateVirtualServer"
var updateVirtualServer VirtualServer

func setup() {
	enabled := false
	newBasicVirtualServer := Basic{
		Enabled:  &enabled,
		Pool:     "pool_test_rui",
		Port:     80,
		Protocol: "http",
	}
	updateVirtualServerProperties := Properties{Basic: newBasicVirtualServer}
	updateVirtualServer = VirtualServer{Properties: updateVirtualServerProperties}

	updateVirtualServerAPI = NewUpdate(updateVirtualServerName, updateVirtualServer)
	updateVirtualServerAPI.SetResponseObject(&updateVirtualServer)
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
	expectedJSON := `{"properties":{"basic":{"enabled":false,"pool":"pool_test_rui","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{},"gzip":{},"ssl":{}}}`
	jsonBytes, err := json.Marshal(updateVirtualServerAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestUpdateGetResponse(t *testing.T) {
	setup()
	getResponse := updateVirtualServerAPI.GetResponse()
	assert.Equal(t, getResponse, updateVirtualServer)

}
