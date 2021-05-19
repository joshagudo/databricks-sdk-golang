package azure_test

import (
	"fmt"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/groups/httpmodels"
	"github.com/stretchr/testify/assert"
)

const (
	pfxGroup = "testGroup"
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

// GroupExists checks if group exists in the list of group names
func GroupExists(groups []string, groupName string) bool {
	for _, g := range groups {
		if g == groupName {
			return true
		}
	}
	return false
}
