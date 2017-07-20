package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

/*
// GetRuleAPI base object.
type GetRuleAPI struct {
	*rest.BaseAPI
}
*/

// NewGetRule : returns a rule
func NewGetRule(ruleName string) *rest.BaseAPI {
	this := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/"+ruleName, nil, new(string), new(api.VTMError))
	return this
}

/*
// GetResponse returns the string representation of the traffic script
func (getRule *GetRuleAPI) GetResponse() string {
	return string(getRule.RawResponse())
}*/
