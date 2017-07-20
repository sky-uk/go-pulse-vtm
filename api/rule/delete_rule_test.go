package rule

import (
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteRuleAPI *rest.BaseAPI
var deleteRuleName string

func setupDelete() {
	deleteRuleName = "test-rule"
	deleteRuleAPI = NewDelete(deleteRuleName)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteRuleAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/"+deleteRuleName, deleteRuleAPI.Endpoint())
}
