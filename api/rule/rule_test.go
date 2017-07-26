package rule

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testRuleEndpoint = "/api/tm/3.8/config/active/rules/"

var createRuleAPI, updateRuleAPI, getAllRuleAPI, getRuleAPI, deleteRuleAPI *rest.BaseAPI
var ruleName = "testRule"
var testAllRulesJSON []byte

func setupRuleTest() {
	trafficScript := []byte(`
	if( string.ipmaskmatch( request.getremoteip(), "192.168.123.10" ) ){
		connection.discard();
	}`)
	createRuleAPI = NewCreate(ruleName, trafficScript)

	getAllRuleAPI = NewGetAll()
	testAllRulesJSON = []byte(`{"children":[{"name":"ruleTestOne","href":"/api/tm/3.8/config/active/rules/ruleTestOne"},{"name":"ruleTestTwo","href":"/api/tm/3.8/config/active/rules/ruleTestTwo"}]}`)

	getRuleAPI = NewGetRule(ruleName)
	updateRuleAPI = NewUpdate(ruleName, trafficScript)
	deleteRuleAPI = NewDelete(ruleName)
}

func TestNewCreateMethod(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, http.MethodPut, createRuleAPI.Method())
}

func TestNewCreateEndpoint(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, testRuleEndpoint+ruleName, createRuleAPI.Endpoint())
}

func TestNewGetAllMethod(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, http.MethodGet, getAllRuleAPI.Method())
}

func TestNewGetAllEndpoint(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, testRuleEndpoint, getAllRuleAPI.Endpoint())
}

func TestNewGetAllUnmarshalling(t *testing.T) {
	setupRuleTest()
	jsonError := json.Unmarshal(testAllRulesJSON, getAllRuleAPI.ResponseObject())

	response := *getAllRuleAPI.ResponseObject().(*Rules)
	assert.Nil(t, jsonError)
	assert.Len(t, response.Children, 2)
	assert.Equal(t, "ruleTestOne", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/rules/ruleTestOne", response.Children[0].HRef)
	assert.Equal(t, "ruleTestTwo", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/rules/ruleTestTwo", response.Children[1].HRef)
}

func TestNewGetRuleMethod(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, http.MethodGet, getRuleAPI.Method())
}

func TestNewGetRuleEndpoint(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, testRuleEndpoint+ruleName, getRuleAPI.Endpoint())
}

func TestNewUpdateMethod(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, http.MethodPut, updateRuleAPI.Method())
}

func TestNewUpdateEndpoint(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, testRuleEndpoint+ruleName, updateRuleAPI.Endpoint())
}

func TestNewDeleteMethod(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, http.MethodDelete, deleteRuleAPI.Method())
}

func TestNewDeleteEndpoint(t *testing.T) {
	setupRuleTest()
	assert.Equal(t, testRuleEndpoint+ruleName, deleteRuleAPI.Endpoint())
}
