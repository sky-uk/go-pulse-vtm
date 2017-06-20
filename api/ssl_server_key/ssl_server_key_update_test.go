package sslServerKey

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateSSLServerKeyAPI *UpdateSSLServerKeyAPI
var updateSSLServerKeyName = "updateSSLServerKey"

func setup() {
	newBasicSSLServerKey := Basic{
		Note:    "test_note",
		Public:  "test_public.com",
		Private: "test_private.com",
		Request: "request",
	}
	newSSLServerKeyProperties := Properties{Basic: newBasicSSLServerKey}
	newSSLServerKey := SSLServerKey{Properties: newSSLServerKeyProperties}

	updateSSLServerKeyAPI = NewUpdate(updateSSLServerKeyName, newSSLServerKey)
	updateSSLServerKeyAPI.SetResponseObject("SSLServerKey updated")
}

func TestUpdateMethod(t *testing.T) {
	setup()
	assert.Equal(t, http.MethodPut, updateSSLServerKeyAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setup()
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/"+updateSSLServerKeyName, updateSSLServerKeyAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	setup()
	expectedJSON := `{"properties":{"basic":{"enabled":false,"pool":"pool_test_rui","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{},"gzip":{},"ssl":{}}}`
	jsonBytes, err := json.Marshal(updateSSLServerKeyAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestUpdateGetResponse(t *testing.T) {
	setup()
	getResponse := updateSSLServerKeyAPI.GetResponse()
	assert.Equal(t, getResponse, "/private/status/check")

}
