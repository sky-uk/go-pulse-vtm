package usergroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

const userGroupsEndpoint = "/api/tm/3.8/config/active/user_groups/"

func NewGet(groupName string) *api.BaseAPI {
	getUserGroupAPI := api.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, nil, new(UserGroup))
	return getUserGroupAPI
}

func NewGetAll() *api.BaseAPI {
	getAllUserGroupAPI := api.NewBaseAPI(http.MethodGet, userGroupsEndpoint, nil, new(UserGroupsList))
	return getAllUserGroupAPI
}

func NewPut(groupName string, userGroup UserGroup) *api.BaseAPI {
	getAllUserGroupAPI := api.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, userGroup, new(UserGroup))
	return getAllUserGroupAPI
}

func NewDelete(groupName string) *api.BaseAPI {
	getAllUserGroupAPI := api.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, nil, nil)
	return getAllUserGroupAPI
}
