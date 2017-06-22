package sslServerKey

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllSSLServerKeysAPI *GetAllSSLServerKeys

func setupGetAll() {
	getAllSSLServerKeysAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllSSLServerKeysAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys", getAllSSLServerKeysAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	jsonContent := []byte("{\"children\":[{\"name\":\"test-ssl-server-key1\",\"href\":\"/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key1\"},{\"name\":\"test-ssl-server-key2\",\"href\":\"/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key2\"}]}")
	jsonErr := json.Unmarshal(jsonContent, getAllSSLServerKeysAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	assert.Len(t, getAllSSLServerKeysAPI.GetResponse().Children, 2)
	assert.Equal(t, "test-ssl-server-key1", getAllSSLServerKeysAPI.GetResponse().Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key1", getAllSSLServerKeysAPI.GetResponse().Children[0].HRef)
	assert.Equal(t, "test-ssl-server-key2", getAllSSLServerKeysAPI.GetResponse().Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key2", getAllSSLServerKeysAPI.GetResponse().Children[1].HRef)
}

func TestGetAllTrafficIPGroupsAPIGetResponse(t *testing.T) {
	setupGetAll()
	assert.IsType(t, getAllSSLServerKeysAPI.GetResponse(), &SSLServerKeysList{})
}
