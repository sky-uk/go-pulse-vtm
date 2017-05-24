package monitor

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/json"
)

var createMonitorAPI *CreateMonitorAPI
var newMonitorName string = "exampleMonitor"

func createSetup() {

	newHTTPMonitor := MonitorHTTP{URIPath: "/download/private/status/check"}
	newBasicMonitor := MonitorBasic{Delay: 6, Failures: 3, Type: "http", Timeout: 4}
	newMonitorProperties := MonitorProperties{Basic: newBasicMonitor, Http: newHTTPMonitor}
	newMonitor := Monitor{Properties: newMonitorProperties}

	createMonitorAPI = NewCreate(newMonitorName, newMonitor)
	createMonitorAPI.SetResponseObject("/download/private/status/check")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createMonitorAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/"+newMonitorName, createMonitorAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := "{\"properties\":{\"basic\":{\"delay\":6,\"failures\":3,\"type\":\"http\",\"timeout\":4},\"http\":{\"path\":\"/download/private/status/check\"}}}"
	jsonBytes, err := json.Marshal(createMonitorAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	getResponse := createMonitorAPI.GetResponse()
	assert.Equal(t, getResponse, "/download/private/status/check")

}

//TODO
func TestCreateUnMarshalling(t *testing.T) {
}