package sslServerKey

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSSLServerKeyAPI *rest.BaseAPI

func setupNewGetSSLServerKey() {
	getSSLServerKeyAPI = NewGet("test-ssl-server-key")
}

func TestNewGetMonitorMethod(t *testing.T) {
	setupNewGetSSLServerKey()
	assert.Equal(t, http.MethodGet, getSSLServerKeyAPI.Method())
}

func TestNewGetMonitorEndpoint(t *testing.T) {
	setupNewGetSSLServerKey()
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key", getSSLServerKeyAPI.Endpoint())
}

func TestNewGetMonitorUnmarshalling(t *testing.T) {
	setupNewGetSSLServerKey()

	jsonContent := []byte("{\"properties\":{\"basic\":{\"note\":\"test\",\"public\":\"test.public.com\",\"private\":\"test.private.com\",\"request\":\"testrequest\"}}}")
	jsonErr := json.Unmarshal(jsonContent, getSSLServerKeyAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	response := *getSSLServerKeyAPI.ResponseObject().(*SSLServerKey)
	assert.Equal(t, "test", response.Properties.Basic.Note)
	assert.Equal(t, "test.public.com", response.Properties.Basic.Public)
	assert.Equal(t, "test.private.com", response.Properties.Basic.Private)
	assert.Equal(t, "testrequest", response.Properties.Basic.Request)
}
