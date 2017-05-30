package pool

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createPool *CreatePoolAPI

func setupCreatePool() {
	node := NewMemberNodes("127.0.0.1:80","active",1 )
	node2 := NewMemberNodes("127.0.0.1:81","active",1 )
	createPool = NewCreate("pool_test_rui_4", []MemberNodes{node,node2}, []string{"ping"})
}