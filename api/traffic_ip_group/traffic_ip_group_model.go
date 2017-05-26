package trafficIpGroups

// TrafficIPGroupList : List of ChildTrafficIPGroup data structure
type TrafficIPGroupList struct {
	Children []ChildTrafficIPGroup `json:"children"`
}

// ChildTrafficIPGroup : Child of TrafficIPGroupList, contains name and href of an Traffic IP Group
type ChildTrafficIPGroup struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}

// TrafficIPGroup : Main traffic IP group data structure
type TrafficIPGroup struct {
	Properties Properties `json:"properties"`
}

// Properties : Stored within TrafficIPGroup, contains Basic data structure
type Properties struct {
	Basic Basic `json:"basic"`
}

// Basic : Stored within Properties, contains attributes of an TrafficIPGroup
type Basic struct {
	Enabled                      bool     `json:"enabled"`
	HashSourcePort               bool     `json:"hash_source_port"`
	IPAssignmentMode             string   `json:"ip_assignment_mode"`
	IPMapping                    []string `json:"ip_mapping"`
	IPAddresses                  []string `json:"ipaddresses"`
	KeepTogether                 bool     `json:"keeptogether"`
	Location                     int      `json:"location"`
	Machines                     []string `json:"machines"`
	Mode                         string   `json:"mode"`
	Multicast                    string   `json:"multicast"`
	Note                         string   `json:"note"`
	RhiBgpMetricBase             int      `json:"rhi_bgp_metric_base"`
	RhiBgpPassiveMetricOffset    int      `json:"rhi_bgp_passive_metric_offset"`
	RhiOspfv2MetricBase          int      `json:"rhi_ospfv2_metric_base"`
	RhiOspfv2PassiveMetricOffset int      `json:"rhi_ospfv2_passive_metric_offset"`
	RhiProtocols                 string   `json:"rhi_protocols"`
	Slaves                       []string `json:"slaves"`
}
