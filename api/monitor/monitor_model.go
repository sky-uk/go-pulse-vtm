package monitor

// Monitor : main Monitor data structure
type Monitor struct {
	Basic Basic `json:"basic"`
	HTTP  HTTP  `json:"http"`
}

// Basic : basic monitor structure
type Basic struct {
	BackOFF   bool   `json:"back_off"`
	Delay     int    `json:"delay"`
	ProtoType string `json:"type"`
}

// HTTP : MonitorHTTP structure
type HTTP struct {
}

// MonitorsList : List of nodes monitored
type MonitorsList struct {
	Children []ChildMonitor `json:"children"`
}

// ChildMonitor : monitored node structure
type ChildMonitor struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}
