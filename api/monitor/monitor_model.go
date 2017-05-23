package monitor

type Monitor struct {
	Basic MonitorBasic `json:"basic"`
	Http  MonitorHTTP  `json:"http"`
}

type MonitorBasic struct {
	BackOFF   bool   `json:"back_off"`
	Delay     int    `json:"delay"`
	ProtoType string `json:"type"`
}
type MonitorHTTP struct {
}

type MonitorsList struct {
	Children []ChildMonitor `json:"children"`
}

type ChildMonitor struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}