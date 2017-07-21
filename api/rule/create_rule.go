package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewCreate : Create a new rule
func NewCreate(ruleName string, trafficScript []byte) *rest.BaseAPI {
	createRuleAPI := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/rules/"+ruleName, trafficScript, nil, new(api.VTMError))
	return createRuleAPI
}
