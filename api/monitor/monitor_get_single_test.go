package monitor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSingleMonitorAPI *GetSingleMonitor

func setupGetSingleMonitor() *MonitorsList {

	var monitorList MonitorsList
	var children = make([]ChildMonitor, 2)

	children[0] = ChildMonitor{Name: "firstmonitor", HRef: "/api/tm/3.8/config/active/monitors/firstmonitor"}
	children[1] = ChildMonitor{Name: "secondmonitor", HRef: "/api/tm/3.8/config/active/monitors/secondmonitor"}
	monitorList.Children = children

	return &monitorList
}

func setupTestMonitorToString() *Monitor {

	monitorHTTP := HTTP{URIPath: "/some/status/page"}
	monitorVerbosity := true
	monitorBasic := Basic{Delay: 7, Failures: 2, Type: "http", Timeout: 11, Verbose: &monitorVerbosity}
	monitorProperties := Properties{Basic: monitorBasic, HTTP: monitorHTTP}
	monitor := Monitor{Properties: monitorProperties}

	return &monitor
}

func setupNewGetMonitor() {
	getSingleMonitorAPI = NewGetSingleMonitor("somemonitor")
}

func TestGetSingleMonitor(t *testing.T) {
	monitorList := setupGetSingleMonitor()

	firstFiltered := monitorList.FilterByName("firstmonitor")
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/firstmonitor", firstFiltered.HRef)

	secondFiltered := monitorList.FilterByName("secondmonitor")
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/secondmonitor", secondFiltered.HRef)
}

func TestMonitorToString(t *testing.T) {
	monitor := setupTestMonitorToString()
	assert.Contains(t, monitor.String(), "/some/status/page")
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
	response := getSingleMonitorAPI.GetResponse()
	assert.Equal(t, 6, response.Properties.Basic.Delay)
	assert.Equal(t, 3, response.Properties.Basic.Failures)
	assert.Equal(t, 4, response.Properties.Basic.Timeout)
	assert.Equal(t, "http", response.Properties.Basic.Type)
	assert.Equal(t, "/download/private/status/check", response.Properties.HTTP.URIPath)
	assert.Equal(t, &verbosityCheck, response.Properties.Basic.Verbose)
}
