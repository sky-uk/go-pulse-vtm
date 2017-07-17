package rule

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// GetAllRules object
type GetAllRules struct {
	*api.BaseAPI
}

// NewGetAll() *GetAllRules {
func NewGetAll() *GetAllRules {
	this := new(GetAllRules)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/rules/", nil, new(RuleList))
	return this
}

// GetResponse returns the response object of GetAllRules
func (getAllRules GetAllRules) GetResponse() *RuleList {
	return getAllRules.ResponseObject().(*RuleList)
}
