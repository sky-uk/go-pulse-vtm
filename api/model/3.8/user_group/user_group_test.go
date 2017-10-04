package userGroup

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUserGroup(t *testing.T) {

	name := api.SetTestResourceName("user_groups_")
	setUserGroup(name, t)
	getUserGroup(name, t)
	deleteUserGroup(name, t)
}

func setUserGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	permissionList := make([]Permission, 0)

	permissionConfigure := Permission{
		Name:        "Configure",
		AccessLevel: "full",
	}
	permissionList = append(permissionList, permissionConfigure)

	permissionGlobalSettings := Permission{
		Name:        "Global_Settings",
		AccessLevel: "full",
	}
	permissionList = append(permissionList, permissionGlobalSettings)

	resource := UserGroup{}
	resource.Properties.Basic.Timeout = 60
	resource.Properties.Basic.Description = "go-brocade-vtm test user group"
	resource.Properties.Basic.PasswordExpireTime = 137
	resource.Properties.Basic.Permissions = permissionList

	newUserGroup := UserGroup{}
	err = client.Set("user_groups", name, resource, &newUserGroup)
	if err != nil {
		t.Fatal("Error creating resource: ", err)
	}
	log.Println("Created User Group ", name)

	assert.Equal(t, uint(60), newUserGroup.Properties.Basic.Timeout)
	assert.Equal(t, "go-brocade-vtm test user group", newUserGroup.Properties.Basic.Description)
	assert.Equal(t, uint(137), newUserGroup.Properties.Basic.PasswordExpireTime)
	assert.Equal(t, 2, len(newUserGroup.Properties.Basic.Permissions))
}

func getUserGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	userGroup := UserGroup{}
	err = client.GetByName("user_groups", name, &userGroup)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found User Group: ", userGroup)

	assert.Equal(t, uint(60), userGroup.Properties.Basic.Timeout)
	assert.Equal(t, "go-brocade-vtm test user group", userGroup.Properties.Basic.Description)
	assert.Equal(t, uint(137), userGroup.Properties.Basic.PasswordExpireTime)
	assert.Equal(t, 2, len(userGroup.Properties.Basic.Permissions))
}

func deleteUserGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("user_groups", name)
	if err != nil {
		t.Fatal("Error deleting resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
