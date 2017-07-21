package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewGetAll : returns a list of rules {
func NewGetAll() *rest.BaseAPI {
	this := rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/", nil, new(Rules), new(api.VTMError))
	return this
}
