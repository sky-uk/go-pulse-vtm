package trafficIpGroups

type TrafficIPGroupList struct {
	Children []ChildTrafficIPGroup `json:"children"`
}

type ChildTrafficIPGroup struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}


type TrafficIPGroup struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Basic Basic `json:"basic"`
}

type Basic struct {
	Enabled bool `json:"enabled"`
	HashSourcePort bool `json:"hash_source_port"`
	IPAssignmentMode string `json:""ip_assignment_mode"`
	IPMapping []string `json:""ip_mapping"`
	IPAddresses []string `json:"ipaddresses"`
	KeepTogether bool `json:"keeptogether"`
	Location int `json:"location"`
	Machines []string `json:"machines"`
	Mode string `json:"mode"`
	Multicast string `json:"multicast"`
	Note string `json:"note"`
	RhiBgpMetricBase int `json:"rhi_bgp_metric_base"`
	RhiBgpPassiveMetricOffset int `json:"rhi_bgp_passive_metric_offset"`
	RhiOspfv2MetricBase int `json:"rhi_ospfv2_metric_base"`
	RhiOspfv2PassiveMetricOffset int `json:"rhi_ospfv2_passive_metric_offset"`
	RhiProtocols string `json:"rhi_protocols"`
	Slaves	[]string `json:"slaves"`
}

