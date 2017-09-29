package dnsZoneFile

// DNSZoneFile : contains the DNS zone file
type DNSZoneFile struct {
	Name   string
	Script string
}

// DNSZoneFiles : List of DNS zone files
type DNSZoneFiles struct {
	Children []ChildRule `json:"children"`
}

// ChildRule : contains a rule's name and href
type ChildRule struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}
