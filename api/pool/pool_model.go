package pool

// Pool - main pool struct
type Pool struct {
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
}

// Properties
type Properties struct {
	Basic          Basic          `json:"basic"`
	HTTP           HTTP           `json:"http"`
	LoadBalancing  LoadBalancing  `json:"load_balancing"`
}

// Basic - main pool definitions
type Basic struct {
	BandwidthClass string     `json:"bandwidth_class"`
	FailurePool    string     `json:"failure_pool"`
	Monitors       []Monitors `json:"children"`
	//NodesTable     []Nodes    `json:"children"`
}


type LoadBalancing struct {
	Algorithm string `json:"algorithm"`
	PriorityEnabled bool `json:"priority_enabled"`
	PriorityNodes   int `json:"priority_nodes"`
}

// Monitor

type Monitors struct {
	MonitorType []string
}

// HTTP
type HTTP struct {
}

type PoolList struct {
	ChildPools []ChildPools `json:"children"`
}

type ChildPools struct {
	Name string `json:"name"`
	Href string `json:"href"`
}
