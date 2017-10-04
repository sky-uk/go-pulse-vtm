package userAuthenticator

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUserAuthenticator(t *testing.T) {

	name := api.SetTestResourceName("user_authenticators")
	setUserAuthenticator(name, t)
	getUserAuthenticator(name, t)
	deleteUserAuthenticator(name, t)
}

func setUserAuthenticator(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	resource := UserAuthenticator{}
	resource.Properties.Basic.Enabled = true
	resource.Properties.Basic.Description = "go-brocade-vtm test user authenticator"
	resource.Properties.Basic.Type = "ldap"

	resource.Properties.LDAP = LDAP{
		BaseDN:         "OU=users",
		DNMethod:       "search",
		SearchPassword: "test-password",
		Filter:         "uid=%u",
		Port:           3890,
		Server:         "10.0.0.100",
		Timeout:        60,
	}

	newUserAuthenticator := UserAuthenticator{}
	err = client.Set("user_authenticators", name, resource, &newUserAuthenticator)
	if err != nil {
		t.Fatal("Error creating resource: ", err)
	}
	log.Println("Created User Authenticator ", name)

	assert.Equal(t, true, newUserAuthenticator.Properties.Basic.Enabled)
	assert.Equal(t, "go-brocade-vtm test user authenticator", newUserAuthenticator.Basic.Description)
	assert.Equal(t, "ldap", newUserAuthenticator.Basic.Type)

	assert.Equal(t, "OU=users", newUserAuthenticator.Properties.LDAP.BaseDN)
	assert.Equal(t, "search", newUserAuthenticator.Properties.LDAP.DNMethod)
	assert.Equal(t, "test-password", newUserAuthenticator.Properties.LDAP.SearchPassword)
	assert.Equal(t, "uid=%u", newUserAuthenticator.Properties.LDAP.Filter)
	assert.Equal(t, uint(3890), newUserAuthenticator.Properties.LDAP.Port)
	assert.Equal(t, "10.0.0.100", newUserAuthenticator.Properties.LDAP.Server)
	assert.Equal(t, uint(60), newUserAuthenticator.Properties.LDAP.Timeout)
}

func getUserAuthenticator(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	userAuthenticator := UserAuthenticator{}
	err = client.GetByName("user_authenticators", name, &userAuthenticator)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found User Authenticator: ", userAuthenticator)

	assert.Equal(t, true, userAuthenticator.Properties.Basic.Enabled)
	assert.Equal(t, "go-brocade-vtm test user authenticator", userAuthenticator.Basic.Description)
	assert.Equal(t, "ldap", userAuthenticator.Basic.Type)

	assert.Equal(t, "OU=users", userAuthenticator.Properties.LDAP.BaseDN)
	assert.Equal(t, "search", userAuthenticator.Properties.LDAP.DNMethod)
	assert.Equal(t, "test-password", userAuthenticator.Properties.LDAP.SearchPassword)
	assert.Equal(t, "uid=%u", userAuthenticator.Properties.LDAP.Filter)
	assert.Equal(t, uint(3890), userAuthenticator.Properties.LDAP.Port)
	assert.Equal(t, "10.0.0.100", userAuthenticator.Properties.LDAP.Server)
	assert.Equal(t, uint(60), userAuthenticator.Properties.LDAP.Timeout)
}

func deleteUserAuthenticator(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("user_authenticators", name)
	if err != nil {
		t.Fatal("Error deleting resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
