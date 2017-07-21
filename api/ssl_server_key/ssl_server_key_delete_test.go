package sslServerKey

import (
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteSSLServerKeyAPI *rest.BaseAPI

func setupDelete() {
	deleteSSLServerKeyAPI = NewDelete("test-ssl-server-key-name")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteSSLServerKeyAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key-name", deleteSSLServerKeyAPI.Endpoint())
}
