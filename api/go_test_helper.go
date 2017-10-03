package api

import (
	"math/rand"
	"time"
	"strconv"
)

// SetTestResourceName : used by resource tests to generate a random unique name.
func SetTestResourceName(prefix string) string {
	rand.Seed(time.Now().Unix())
	return prefix + strconv.Itoa(rand.Int())
}
