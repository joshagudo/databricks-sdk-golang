package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/groups/httpmodels"
)

// GroupsAPI exposes the Groups API
type GroupsAPI struct {
	Client DBClient
}

func (a GroupsAPI) init(client DBClient) GroupsAPI {
	a.Client = client
	return a
}

// AddMember adds a user or group to a group
func (a GroupsAPI) AddMember(addMemberReq httpmodels.AddMemberReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/add-member", addMemberReq, nil)
	return err
}

// Create creates a new group with the given name
func (a GroupsAPI) Create(createReq httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResponse httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/groups/create", createReq, nil)
	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// ListMembers returns all of the members of a particular group
func (a GroupsAPI) ListMembers(listMembersReq httpmodels.ListMembersReq) (httpmodels.ListMembersResp, error) {
	var listMembersResp httpmodels.ListMembersResp

	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list-members", listMembersReq, nil)
	if err != nil {
		return listMembersResp, err
	}

	err = json.Unmarshal(resp, &listMembersResp)
	return listMembersResp, err
}

// List returns all of the groups in an organization
func (a GroupsAPI) List() (httpmodels.ListResp, error) {
	var listResponse httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list", nil, nil)
	if err != nil {
		return listResponse, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse, err
}

// ListParents retrieves all groups in which a given user or group is a member
func (a GroupsAPI) ListParents(listParentsReq httpmodels.ListParentsReq) (httpmodels.ListParentsResp, error) {
	var listParentsResponse httpmodels.ListParentsResp

	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list-parents", listParentsReq, nil)
	if err != nil {
		return listParentsResponse, err
	}

	err = json.Unmarshal(resp, &listParentsResponse)
	return listParentsResponse, err
}

// RemoveMember removes a user or group from a group
func (a GroupsAPI) RemoveMember(removeMemberReq httpmodels.RemoveMemberReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/remove-member", removeMemberReq, nil)
	return err
}

// Delete removes a group from this organization
func (a GroupsAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/delete", deleteReq, nil)
	return err
}
