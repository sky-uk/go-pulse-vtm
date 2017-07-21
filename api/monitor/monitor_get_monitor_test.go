package monitor

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSingleMonitorAPI *rest.BaseAPI

func setupNewGetMonitor() {
	getSingleMonitorAPI = NewGet("somemonitor")
}

func TestNewGetMonitorMethod(t *testing.T) {
	setupNewGetMonitor()
	assert.Equal(t, http.MethodGet, getSingleMonitorAPI.Method())
}

func TestNewGetMonitorEndpoint(t *testing.T) {
	setupNewGetMonitor()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/somemonitor", getSingleMonitorAPI.Endpoint())
}

func TestNewGetMonitorUnmarshalling(t *testing.T) {
	setupNewGetMonitor()
	verbosityCheck := true

	jsonContent := []byte("{\"properties\":{\"basic\":{\"delay\":6,\"failures\":3,\"timeout\":4,\"type\":\"http\",\"verbose\":true},\"http\":{\"path\":\"/download/private/status/check\"}}}")
	jsonErr := json.Unmarshal(jsonContent, getSingleMonitorAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	response := getSingleMonitorAPI.ResponseObject().(*Monitor)
	assert.Equal(t, uint(6), response.Properties.Basic.Delay)
	assert.Equal(t, uint(3), response.Properties.Basic.Failures)
	assert.Equal(t, uint(4), response.Properties.Basic.Timeout)
	assert.Equal(t, "http", response.Properties.Basic.Type)
	assert.Equal(t, "/download/private/status/check", response.Properties.HTTP.URIPath)
	assert.Equal(t, &verbosityCheck, response.Properties.Basic.Verbose)
}
