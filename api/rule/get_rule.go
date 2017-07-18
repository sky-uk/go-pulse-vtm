package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// GetRule base object.
type GetRule struct {
	*api.BaseAPI
}

// NewGetRule : returns a rule
func NewGetRule(ruleName string) *GetRule {
	this := new(GetRule)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/"+ruleName, nil, new(string))
	return this
}

// GetResponse returns the string representation of the traffic script
func (getRule *GetRule) GetResponse() string {
	return getRule.ResponseObject().(string)
}

/*
// GetResponse : get response object from created zone
func (cza *CreateZoneAuthAPI) GetResponse() string {
	return *cza.ResponseObject().(*string)
}

*/
