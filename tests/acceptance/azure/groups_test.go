package azure_test

import (
	"fmt"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/groups/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/groups/models"
	"github.com/stretchr/testify/assert"
)

const (
	pfxGroup    = "testGroup"
	adminsGroup = "admins"
)

func TestAzureGroupsCreateListDelete(t *testing.T) {

	groupName := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	// Test create group and assert no errors
	createReq := httpmodels.CreateReq{
		GroupName: groupName,
	}
	createRes, e := c.Groups().Create(createReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create group %+v", createReq))
	assert.NotEmpty(t, createRes.GroupName, fmt.Sprintf("failed to create group %+v", createReq))

	// Test list groups and assert group exists
	listRes, e := c.Groups().List()
	assert.Nil(t, e, "failed to list groups")
	assert.True(t, GroupExists(listRes.GroupNames, groupName), fmt.Sprintf("created group not found %+v", createReq))

	// Test delete group and assert no errors
	deleteReq := httpmodels.DeleteReq{
		GroupName: groupName,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete group %+v", deleteReq))

	// Test list groups and assert group does not exist
	listRes, e = c.Groups().List()
	assert.Nil(t, e, "failed to list groups")
	assert.False(t, GroupExists(listRes.GroupNames, groupName), fmt.Sprintf("deleted group still exists %+v", deleteReq))
}

func TestAzureGroupMembersAddListRemove(t *testing.T) {

	// Setup: List members in a default 'users' group, assess no errors and pick one user for further tests purposes
	listAdminReq := httpmodels.ListMembersReq{
		GroupName: adminsGroup,
	}
	listAdminResp, e := c.Groups().ListMembers(listAdminReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", adminsGroup))

	hasAdmins := listAdminResp.Members != nil && len(*listAdminResp.Members) > 0
	assert.True(t, e != nil || hasAdmins, fmt.Sprintf("no members in a group %s", adminsGroup))
	if !hasAdmins {
		return
	}

	userName := (*listAdminResp.Members)[0].UserName

	// Setup: create group
	groupName := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	createGroupReq := httpmodels.CreateReq{
		GroupName: groupName,
	}
	createRes, e := c.Groups().Create(createGroupReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create group %+v", createGroupReq))
	assert.NotEmpty(t, createRes.GroupName, fmt.Sprintf("failed to create group %+v", createGroupReq))

	// Add member to a group and assert no errors
	addMemberReq := httpmodels.AddMemberReq{
		UserName:   userName,
		ParentName: groupName,
	}
	e = c.Groups().AddMember(addMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to add member to a group %+v", addMemberReq))

	// List members in a group and assert member added
	listMemberReq := httpmodels.ListMembersReq{
		GroupName: groupName,
	}
	listMemberResp, e := c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", groupName))
	assert.True(t, GroupMemberExists(listMemberResp.Members, userName), fmt.Sprintf("added member not found in a group %+v", addMemberReq))

	// Remove member from a group and assert no errors
	removeMemberReq := httpmodels.RemoveMemberReq{
		UserName:   userName,
		ParentName: groupName,
	}
	e = c.Groups().RemoveMember(removeMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to remove member from a group %+v", removeMemberReq))

	// List members in a group and assert member removed
	listMemberResp, e = c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", groupName))
	assert.False(t, GroupMemberExists(listMemberResp.Members, userName), fmt.Sprintf("removed member still exists in a group %+v", removeMemberReq))

	// Cleanup: delete group
	deleteReq := httpmodels.DeleteReq{
		GroupName: groupName,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete group %+v", deleteReq))
}

// GroupExists checks if group exists in the list of group names
func GroupExists(groups []string, groupName string) bool {
	for _, g := range groups {
		if g == groupName {
			return true
		}
	}
	return false
}

// GroupMemberExists checks if user exists in the list of group members
func GroupMemberExists(membersPtr *[]models.PrincipalName, userName string) bool {
	if membersPtr != nil {
		for _, m := range *membersPtr {
			if m.UserName == userName {
				return true
			}
		}
	}
	return false
}
