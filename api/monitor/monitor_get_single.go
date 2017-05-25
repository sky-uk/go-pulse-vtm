package monitor

import (
	"fmt"
)

// FilterByName returns a monitor object if the monitor name matches.
func (monitorList MonitorsList) FilterByName(name string) *ChildMonitor {
	var foundMonitor ChildMonitor
	for _, childMonitor := range monitorList.Children {
		if childMonitor.Name == name {
			foundMonitor = childMonitor
			break
		}
	}
	return &foundMonitor
}

// String returns a string representation of the monitor
func (monitor Monitor) String() string {
	return fmt.Sprintf("Monitor: %s", monitor.Properties)
}
