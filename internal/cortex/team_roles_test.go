package cortex_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cortexapps/terraform-provider-cortex/internal/cortex"
	"github.com/stretchr/testify/suite"
)

/** Suite *************************************************************************************************************/

type teamRolesSuite struct {
	suite.Suite
}

func TestCortexTeamRoles(t *testing.T) {
	suite.Run(t, new(teamRolesSuite))
}

func (s *teamRolesSuite) newTeamRole() *cortex.TeamRole {
	return &cortex.TeamRole{
		ID:                   1,
		Tag:                  "engineer",
		Name:                 "Engineer",
		Description:          "A team role",
		NotificationsEnabled: true,
	}
}

func (s *teamRolesSuite) TestGet() {
	id := int64(1)
	expectedResp := s.newTeamRole()
	ctx := context.Background()
	c, teardown, err := setupClient(cortex.Route("teams_roles", fmt.Sprintf("%d", id)), expectedResp, AssertRequestMethod(s.T(), "GET"))
	defer teardown()
	s.NoError(err)

	actualResp, err := c.TeamRoles().Get(ctx, id)
	s.NoError(err)
	s.Equal(expectedResp, actualResp)
}

func (s *teamRolesSuite) TestList() {
	expectedResp := &cortex.TeamRolesListResponse{
		Roles: []cortex.TeamRole{
			*s.newTeamRole(),
		},
	}
	ctx := context.Background()
	c, teardown, err := setupClient(cortex.Route("teams_roles", ""), expectedResp, AssertRequestMethod(s.T(), "GET"))
	defer teardown()
	s.NoError(err)

	actualResp, err := c.TeamRoles().List(ctx, &cortex.TeamRolesListParams{})
	s.NoError(err)
	s.Equal(expectedResp, actualResp)
}

func (s *teamRolesSuite) TestCreate() {
	req := cortex.CreateTeamRoleRequest{
		Name:                 "Engineer",
		Tag:                  "engineer",
		Description:          "A team role",
		NotificationsEnabled: true,
	}
	expectedResp := s.newTeamRole()
	ctx := context.Background()
	c, teardown, err := setupClient(
		cortex.Route("teams_roles", ""),
		expectedResp,
		AssertRequestMethod(s.T(), "POST"),
		AssertRequestBody(s.T(), req),
	)
	defer teardown()
	s.NoError(err)

	actualResp, err := c.TeamRoles().Create(ctx, req)
	s.NoError(err)
	s.Equal(expectedResp, actualResp)
}

func (s *teamRolesSuite) TestUpdate() {
	id := int64(1)
	req := cortex.UpdateTeamRoleRequest{
		ID:                   id,
		Name:                 "Engineer",
		Description:          "A team role",
		NotificationsEnabled: true,
	}
	expectedResp := s.newTeamRole()
	ctx := context.Background()
	c, teardown, err := setupClient(
		cortex.Route("teams_roles", fmt.Sprintf("%d", id)),
		expectedResp,
		AssertRequestMethod(s.T(), "PUT"),
		AssertRequestBody(s.T(), req),
	)
	defer teardown()
	s.NoError(err)

	actualResp, err := c.TeamRoles().Update(ctx, req)
	s.NoError(err)
	s.Equal(expectedResp, actualResp)
}

func (s *teamRolesSuite) TestDelete() {
	id := int64(2)
	ctx := context.Background()
	c, teardown, err := setupClient(cortex.Route("teams_roles", fmt.Sprintf("%d", id)), nil, AssertRequestMethod(s.T(), "DELETE"))
	defer teardown()
	s.NoError(err)

	err = c.TeamRoles().Delete(ctx, id)
	s.NoError(err)
}
