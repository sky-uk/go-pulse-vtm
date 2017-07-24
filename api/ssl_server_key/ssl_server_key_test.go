package sslServerKey

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testSSlServerKeyEndpoint = "/api/tm/3.8/config/active/ssl/server_keys"

var createSSLServerKeyAPI, getAllSSLServerKeyAPI, getSSLServerKeyAPI, updateSSLServerKeyAPI, deleteSSLServerKeyAPI *rest.BaseAPI
var testSSLServerKeyName, marshallingTestExpectedJSON string
var testSSLServerKey SSLServerKey
var testGetAllUnmarshalTestJSON, testGetUnmarshalTestJSON []byte

func setupSSLServerKeyTest() {
	testSSLServerKeyName = "exampleSSLServerKey"
	testSSLServerKeyBasic := Basic{Note: "test", Private: "testprivate.com", Public: "testpublic.com", Request: "test"}
	testSSLServerKeyProperties := Properties{Basic: testSSLServerKeyBasic}
	testSSLServerKey = SSLServerKey{Properties: testSSLServerKeyProperties}
	marshallingTestExpectedJSON = `{"properties":{"basic":{"note":"test","private":"testprivate.com","public":"testpublic.com","request":"test"}}}`

	createSSLServerKeyAPI = NewCreate(testSSLServerKeyName, testSSLServerKey)
	createSSLServerKeyAPI.SetResponseObject(&testSSLServerKey)

	getAllSSLServerKeyAPI = NewGetAll()
	testGetAllUnmarshalTestJSON = []byte(`{"children":[{"name":"test-ssl-server-key1","href":"/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key1"},{"name":"test-ssl-server-key2","href":"/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key2"}]}`)

	getSSLServerKeyAPI = NewGet(testSSLServerKeyName)
	testGetUnmarshalTestJSON = []byte(`{"properties":{"basic":{"note":"test","public":"test.public.com","private":"test.private.com","request":"testrequest"}}}`)

	updateSSLServerKeyAPI = NewUpdate(testSSLServerKeyName, testSSLServerKey)
	updateSSLServerKeyAPI.SetResponseObject(&testSSLServerKey)

	deleteSSLServerKeyAPI = NewDelete(testSSLServerKeyName)
}

func TestNewCreateMethod(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, http.MethodPut, createSSLServerKeyAPI.Method())
}

func TestNewCreateEndpoint(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, testSSlServerKeyEndpoint+"/"+testSSLServerKeyName, createSSLServerKeyAPI.Endpoint())
}

func TestNewCreateMarshalling(t *testing.T) {
	setupSSLServerKeyTest()
	jsonBytes, err := json.Marshal(createSSLServerKeyAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewGetAllMethod(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, http.MethodGet, getAllSSLServerKeyAPI.Method())
}

func TestNewGetAllEndpoint(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, testSSlServerKeyEndpoint, getAllSSLServerKeyAPI.Endpoint())
}

func TestNewGetAllUnmarshalling(t *testing.T) {
	setupSSLServerKeyTest()
	jsonErr := json.Unmarshal(testGetAllUnmarshalTestJSON, getAllSSLServerKeyAPI.ResponseObject())
	response := *getAllSSLServerKeyAPI.ResponseObject().(*SSLServerKeysList)

	assert.Nil(t, jsonErr)
	assert.Len(t, response.Children, 2)
	assert.Equal(t, "test-ssl-server-key1", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key1", response.Children[0].HRef)
	assert.Equal(t, "test-ssl-server-key2", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/ssl/server_keys/test-ssl-server-key2", response.Children[1].HRef)
}

func TestNewGetMethod(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, http.MethodGet, getSSLServerKeyAPI.Method())
}

func TestNewGetEndpoint(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, testSSlServerKeyEndpoint+"/"+testSSLServerKeyName, getSSLServerKeyAPI.Endpoint())
}

func TestNewGetUnmarshalling(t *testing.T) {
	setupSSLServerKeyTest()
	jsonErr := json.Unmarshal(testGetUnmarshalTestJSON, getSSLServerKeyAPI.ResponseObject())
	response := *getSSLServerKeyAPI.ResponseObject().(*SSLServerKey)

	assert.Nil(t, jsonErr)
	assert.Equal(t, "test", response.Properties.Basic.Note)
	assert.Equal(t, "test.public.com", response.Properties.Basic.Public)
	assert.Equal(t, "test.private.com", response.Properties.Basic.Private)
	assert.Equal(t, "testrequest", response.Properties.Basic.Request)
}

func TestNewUpdateMethod(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, http.MethodPut, updateSSLServerKeyAPI.Method())
}

func TestNewUpdateEndpoint(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, testSSlServerKeyEndpoint+"/"+testSSLServerKeyName, updateSSLServerKeyAPI.Endpoint())
}

func TestNewUpdateMarshalling(t *testing.T) {
	setupSSLServerKeyTest()
	jsonBytes, err := json.Marshal(updateSSLServerKeyAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewDeleteMethod(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, http.MethodDelete, deleteSSLServerKeyAPI.Method())
}

func TestNewDeleteEndpoint(t *testing.T) {
	setupSSLServerKeyTest()
	assert.Equal(t, testSSlServerKeyEndpoint+"/"+testSSLServerKeyName, deleteSSLServerKeyAPI.Endpoint())
}
