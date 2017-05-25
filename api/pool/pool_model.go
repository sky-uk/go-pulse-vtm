package pool

// Pool - main pool struct
type Pool struct {
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
}

type SinglePool struct {
	Properties Properties `json:"properties"`
}

// Properties
type Properties struct {
	Basic         Basic         `json:"basic"`
	Connection    Connection    `json:"connection"`
	HTTP          HTTP          `json:"http"`
	LoadBalancing LoadBalancing `json:"load_balancing"`
	Node          Node          `json:"node"`
	Ssl           Ssl           `json:"ssl"`
}

// Basic - main pool definitions
type Basic struct {
	BandwidthClass               string        `json:"bandwidth_class"`
	FailurePool                  string        `json:"failure_pool"`
	MaxConnectionAttempts        int           `json:"max_connection_attempts"`
	MaxIdleConnectionsPerNode    int           `json:"max_idle_connections_pernode"`
	MaxTimeoutConnectionAttempts int           `json:"max_timed_out_connection_attempts"`
	Monitors                     []string      `json:"monitors"`
	NodeCloseWithReset           bool          `json:"node_close_with_rst"`
	NodeConnectionAttempts       int           `json:"node_connection_attempts"`
	NodeDeleteBehavior           string        `json:"node_delete_behavior"`
	NodeDrainDeleteTimeout       int           `json:"node_drain_to_delete_timeout"`
	NodesTable                   []MemberNodes `json:"nodes_table"`
	Note                         string        `json:"note"`
	PassiveMonitoring            bool          `json:"passive_monitoring"`
	PersistentClass              string        `json:"persistent_class"`
	Transparent                  bool          `json:"transparent"`
}

type Connection struct {
	MaxConnectTime        int `json:"max_connect_time"`
	MaxConnectionsPerNode int `json:"max_connections_per_node"`
	MaxQueueSize          int `json:"max_queue_size"`
	MaxReplyTime          int `json:"max_reply_time"`
	QueueTimeout          int `json:"queue_timeout"`
}

// HTTP
type HTTP struct {
	HttpKeepAlive              bool `json:"keepalive"`
	HttpKeepAliveNonIdempotent bool `json:"keepalive_non_idempotent"`
}

// LoadBalancing
type LoadBalancing struct {
	Algorithm       string `json:"algorithm"`
	PriorityEnabled bool   `json:"priority_enabled"`
	PriorityNodes   int    `json:"priority_nodes"`
}

type Node struct {
	CloseOnDeath  bool `json:"close_on_death"`
	RetryFailTime int  `json:"retry_fail_time"`
}

type Ssl struct {
	ClientAuth      bool     `json:"client_auth"`
	CommonNameMatch []string `json:"common_name_match"`
	ElipticCurves   []string `json:"eliptic_curves"`
	Enabled         bool     `json:"enabled"`
	Enhance         bool     `json:"enhance"`
}

type MemberNodes struct {
	Node   string `json:"node"`
	State  string `json:"state"`
	Weight int    `json:"weight"`
}

type PoolList struct {
	ChildPools []ChildPools `json:"children"`
}

type ChildPools struct {
	Name string `json:"name"`
	Href string `json:"href"`
}
