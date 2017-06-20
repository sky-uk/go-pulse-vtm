package sslServerKey

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createSSLServerKeyAPI *CreateSSLServerKeyAPI
var newSSLServerKeyName = "exampleSSLServerKey"

func createSetup() {
	newSSLServerKeyBasic := Basic{Note: "test", Private: "testprivate.com", Public: "testpublic.com", Request: "test"}
	newSSLServerKeyProperties := Properties{Basic: newSSLServerKeyBasic}
	newSSLServerKey := SSLServerKey{Properties: newSSLServerKeyProperties}

	createSSLServerKeyAPI = NewCreate(newSSLServerKeyName, &newSSLServerKey)
	createSSLServerKeyAPI.SetResponseObject("test response")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createSSLServerKeyAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/"+ newSSLServerKeyName, createSSLServerKeyAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := "{\"properties\":{\"basic\":{\"note\":\"test\",\"private\":\"testprivate.com\",\"public\":\"testpublic.com\",\"request\":\"test\"}}}"
	jsonBytes, err := json.Marshal(createSSLServerKeyAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	getResponse := createSSLServerKeyAPI.GetResponse()
	assert.Equal(t, getResponse, "test response")
}
