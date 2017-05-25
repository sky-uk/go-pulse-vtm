package monitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (*MonitorsList, Monitor) {

	var monitorList MonitorsList
	var children = make([]ChildMonitor, 2)

	children[0] = ChildMonitor{Name: "firstmonitor", HRef: "/api/tm/3.8/config/active/monitors/firstmonitor"}
	children[1] = ChildMonitor{Name: "secondmonitor", HRef: "/api/tm/3.8/config/active/monitors/secondmonitor"}

	monitorList.Children = children

	monitorHTTP := HTTP{URIPath: "/some/status/page"}
	monitorBasic := Basic{Delay: 7, Failures: 2, Type: "http", Timeout: 11}
	monitorProperties := Properties{Basic: monitorBasic, HTTP: monitorHTTP}
	monitor := Monitor{Properties: monitorProperties}

	return &monitorList, monitor
}

func TestGetSingleMonitor(t *testing.T) {
	monitorList, _ := setup()

	firstFiltered := monitorList.FilterByName("firstmonitor")
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/firstmonitor", firstFiltered.HRef)

	secondFiltered := monitorList.FilterByName("secondmonitor")
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/secondmonitor", secondFiltered.HRef)
}

func TestMonitorToString(t *testing.T) {
	_, monitor := setup()
	assert.Contains(t, monitor.String(), "/some/status/page")
}
