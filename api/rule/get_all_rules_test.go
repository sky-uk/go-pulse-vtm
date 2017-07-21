package rule

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllRulesAPI *rest.BaseAPI

func setupGetAll() {
	getAllRulesAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllRulesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/", getAllRulesAPI.Endpoint())
}

func TestGetAllUnmarshalling(t *testing.T) {
	setupGetAll()
	jsonRules := []byte(`{"children":[{"name":"ruleTestOne","href":"/api/tm/3.8/config/active/rules/ruleTestOne"},{"name":"ruleTestTwo","href":"/api/tm/3.8/config/active/rules/ruleTestTwo"}]}`)
	jsonError := json.Unmarshal(jsonRules, getAllRulesAPI.ResponseObject())

	response := getAllRulesAPI.ResponseObject().(*Rules)
	assert.Nil(t, jsonError)
	assert.Len(t, response.Children, 2)
	assert.Equal(t, "ruleTestOne", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/rules/ruleTestOne", response.Children[0].HRef)
	assert.Equal(t, "ruleTestTwo", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/rules/ruleTestTwo", response.Children[1].HRef)
}
