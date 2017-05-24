package monitor

// Monitor : main Monitor data structure
type Monitor struct {
	Properties Properties `json:"properties"`
}

// Properties : Properties contains the overall monitor configuration
type Properties struct {
	Basic Basic `json:"basic"`
	HTTP  HTTP  `json:"http"`
}

// Basic : Basic monitor configration
type Basic struct {
	Delay    int    `json:"delay"`
	Failures int    `json:"failures"`
	Type     string `json:"type"`
	Timeout  int    `json:"timeout"`
	UseSSL   bool   `json:"use_ssl,omitempty"`
	Verbose  bool   `json:"verbose,omitempty"`
}

// HTTP : HTTP monitor set up
type HTTP struct {
	Authentication string `json:"authentication,omitempty"`
	BodyRegex      string `json:"body_regex,omitempty"`
	HostHeader     string `json:"host_header,omitempty"`
	URIPath        string `json:"path"`
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
