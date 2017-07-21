package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewGetRule : returns a rule
func NewGetRule(ruleName string) *rest.BaseAPI {
	getRuleAPI := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/"+ruleName, nil, new(string), new(api.VTMError))
	return getRuleAPI
}
