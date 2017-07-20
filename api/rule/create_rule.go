package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

/*
// CreateRuleAPI : create rule API object
type CreateRuleAPI struct {
	*rest.BaseAPI
}
*/

// NewCreate : Create a new rule
func NewCreate(ruleName string, trafficScript []byte) *rest.BaseAPI {
	//this := new(rest.BaseAPI)
	this := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/rules/"+ruleName, trafficScript, nil, new(api.VTMError))
	return this
}
