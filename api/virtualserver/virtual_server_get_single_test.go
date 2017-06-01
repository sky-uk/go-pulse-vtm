package virtualserver

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSingleVirtualServerAPI *GetSingleVirtualServer

func setupGetSingle() {
	getSingleVirtualServerAPI = NewGetSingle("PaaSExampleHTTPvirtualserver")
}

func TestGetSingleMethod(t *testing.T) {
	setupGetSingle()
	assert.Equal(t, http.MethodGet, getSingleVirtualServerAPI.Method())
}

func TestGetSingleEndpoint(t *testing.T) {
	setupGetSingle()
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver", getSingleVirtualServerAPI.Endpoint())
}

func TestGetSingleUnMarshalling(t *testing.T) {
	setupGetSingle()
	fmt.Println(getSingleVirtualServerAPI.GetResponse().Properties)
	jsonContent := `{"properties":{"basic":{"enabled":false,"pool":"pool_test_rui","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{"ssl_data":false},"gzip":{},"ssl":{}}}`
	jsonErr := json.Unmarshal([]byte(jsonContent), getSingleVirtualServerAPI.ResponseObject())
	assert.Nil(t, jsonErr)
}
