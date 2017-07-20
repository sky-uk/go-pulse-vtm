package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

/*
// DeleteRuleAPI : object used to call delete on a rule
type DeleteRuleAPI struct {
	*rest.BaseAPI
}
*/

// NewDelete : used to delete a rule
func NewDelete(ruleName string) *rest.BaseAPI {
	//this := new(rest.BaseAPI)
	this := rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/rules/"+ruleName, nil, nil, new(api.VTMError))
	return this
}
