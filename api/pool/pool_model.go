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
	Load_Balancing Load_Balancing `json:"load_balancing"`
}

// Basic - main pool definitions
type Basic struct {
	Monitors    []Monitors `json:"children"`
	Nodes_Table []Nodes    `json:"children"`
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
