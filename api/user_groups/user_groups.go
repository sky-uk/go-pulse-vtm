package usergroups

import (
	"github.com/sky-uk/go-rest-api"
	"net/http"
	"github.com/sky-uk/go-brocade-vtm/api"
)

const userGroupsEndpoint = "/api/tm/3.8/config/active/user_groups/"

func NewGet(groupName string) *rest.BaseAPI  {
	getUserGroupAPI := rest.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, nil, new(UserGroup), new(api.VTMError))
	return getUserGroupAPI
}

func NewGetAll() *rest.BaseAPI  {
	getAllUserGroupAPI := rest.NewBaseAPI(http.MethodGet, userGroupsEndpoint, nil, new(UserGroupsList), new(api.VTMError))
	return getAllUserGroupAPI
}

func NewPut(groupName string, userGroup UserGroup) *rest.BaseAPI {
	getAllUserGroupAPI := rest.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, userGroup, new(UserGroup), new(api.VTMError))
	return getAllUserGroupAPI
}

func NewDelete(groupName string) *rest.BaseAPI  {
	getAllUserGroupAPI := rest.NewBaseAPI(http.MethodGet, userGroupsEndpoint+groupName, nil, nil, new(api.VTMError))
	return getAllUserGroupAPI
}
