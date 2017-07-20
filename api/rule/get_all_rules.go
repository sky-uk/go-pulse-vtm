package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

/*
// GetAllRules : used for retrieving a list of rules and their href
type GetAllRules struct {
	*rest.BaseAPI
}
*/

// NewGetAll : returns a list of rules {
func NewGetAll() *rest.BaseAPI {
	this := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/", nil, new(Rules), new(api.VTMError))
	return this
}

/*
// GetResponse returns the response object of GetAllRules
func (getAllRules rest.BaseAPI) GetResponse() *Rules {
	return getAllRules.ResponseObject().(*Rules)
}
*/
