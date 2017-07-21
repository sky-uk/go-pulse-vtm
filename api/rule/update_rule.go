package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewUpdate : Update a rule
func NewUpdate(ruleName string, trafficScript []byte) *rest.BaseAPI {
	updateRuleAPI := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/rules/"+ruleName, trafficScript, nil, new(api.VTMError))
	return updateRuleAPI
}
