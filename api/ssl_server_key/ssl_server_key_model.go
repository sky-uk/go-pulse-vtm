package sslServerKey

// A SSLServerKey is a Brocade trusted certificate.
type SSLServerKey struct {
	Properties Properties `json:"properties"`
}

// Properties : Properties contains the overall SSLServerKey configuration
type Properties struct {
	Basic Basic `json:"basic,omitempty"`
}

// Basic struct for SSlServerKey
type Basic struct {
	Note    string `json:"note,omitempty"`
	Private string `json:"private,omitempty"`
	Public  string `json:"public,omitempty"`
	Request string `json:"request,omitempty"`
}

// MonitorsList : List of nodes monitored
type SSLServerKeysList struct {
	Children []ChildSSLServerKey `json:"children"`
}

// ChildMonitor : monitored node structure
type ChildSSLServerKey struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}