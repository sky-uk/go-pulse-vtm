package rule

// RuleList : List of rules
type RuleList struct {
	Children []ChildRule `json:"children"`
}

// ChildRule : contains a rule's name and href
type ChildRule struct {
	Name string `json:"name"`
	HRef string `json:"href"`
}
