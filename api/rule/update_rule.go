package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

/*
// UpdateRuleAPI : update rule API object
type UpdateRuleAPI struct {
	*rest.BaseAPI
}*/

// NewUpdate : Update a rule
func NewUpdate(ruleName string, trafficScript []byte) *rest.BaseAPI {
	this := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/rules/"+ruleName, trafficScript, nil, new(api.VTMError))
	return this
}
