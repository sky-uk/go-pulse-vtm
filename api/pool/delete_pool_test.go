package pool

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var DeleteAPI *DeletePool

func setupDelete() {
	DeleteAPI = NewDelete("pool_test_rui_2")
}

func TestNewDelete(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, DeleteAPI.Method())
}
