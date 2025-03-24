package cortex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dghubble/sling"
)

type TeamRolesClientInterface interface {
	Get(ctx context.Context, id int64) (*TeamRole, error)
	List(ctx context.Context, params *TeamRolesListParams) (*TeamRolesListResponse, error)
	Create(ctx context.Context, req CreateTeamRoleRequest) (*TeamRole, error)
	Update(ctx context.Context, req UpdateTeamRoleRequest) (*TeamRole, error)
	Delete(ctx context.Context, id int64) error
}

type TeamRolesClient struct {
	client *HttpClient
}

var _ TeamRolesClientInterface = &TeamRolesClient{}

func (c *TeamRolesClient) Client() *sling.Sling {
	return c.client.Client()
}

/***********************************************************************************************************************
 * Types
 **********************************************************************************************************************/

type TeamRole struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	Tag                  string `json:"tag"`
	Description          string `json:"description,omitempty"`
	NotificationsEnabled bool   `json:"notificationsEnabled,omitempty"`
}

/***********************************************************************************************************************
 * GET /api/v1/teams/roles/:id
 **********************************************************************************************************************/

func (c *TeamRolesClient) Get(_ context.Context, id int64) (*TeamRole, error) {
	apiResponse := &TeamRole{}
	apiError := &ApiError{}
	response, err := c.Client().Get(Route("teams_roles", fmt.Sprintf("%d", id))).Receive(apiResponse, apiError)
	if err != nil {
		return apiResponse, errors.New("could not get team role: " + err.Error())
	}

	err = c.client.handleResponseStatus(response, apiError)
	if err != nil {
		return apiResponse, errors.Join(errors.New("Failed getting team role: "), err)
	}

	return apiResponse, nil
}

/***********************************************************************************************************************
 * GET /api/v1/teams/roles
 **********************************************************************************************************************/

// TeamRolesListParams are the query parameters for the GET /v1/teams/roles endpoint.
type TeamRolesListParams struct {
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"size,omitempty"`
	Query   string `url:"query,omitempty"`
}

// TeamRolesListResponse is the response from the GET /v1/teams/roles endpoint.
type TeamRolesListResponse struct {
	Roles []TeamRole `json:"items"`
}

func (c *TeamRolesClient) List(_ context.Context, params *TeamRolesListParams) (*TeamRolesListResponse, error) {
	apiResponse := &TeamRolesListResponse{}
	apiError := &ApiError{}

	response, err := c.Client().Get(Route("teams_roles", "")).QueryStruct(&params).Receive(apiResponse, apiError)
	if err != nil {
		return nil, errors.New("could not get team roles: " + err.Error())
	}

	err = c.client.handleResponseStatus(response, apiError)
	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}

/***********************************************************************************************************************
 * POST /api/v1/teams/roles
 **********************************************************************************************************************/

type CreateTeamRoleRequest struct {
	Tag                  string `json:"tag"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	NotificationsEnabled bool   `json:"notificationsEnabled,omitempty"`
}

// ToCreateRequest converts a TeamRole to a CreateTeamRoleRequest.
func (r *TeamRole) ToCreateRequest() CreateTeamRoleRequest {
	return CreateTeamRoleRequest{
		Tag:                  r.Tag,
		Name:                 r.Name,
		Description:          r.Description,
		NotificationsEnabled: r.NotificationsEnabled,
	}
}

func (c *TeamRolesClient) Create(_ context.Context, req CreateTeamRoleRequest) (*TeamRole, error) {
	apiResponse := &TeamRole{}
	apiError := &ApiError{}

	response, err := c.Client().Post(Route("teams_roles", "")).BodyJSON(&req).Receive(apiResponse, apiError)
	if err != nil {
		return apiResponse, errors.New("could not create team role: " + err.Error())
	}

	err = c.client.handleResponseStatus(response, apiError)
	if err != nil {
		reqJson, _ := json.Marshal(req)
		log.Printf("Failed creating team role: %+v\n\nRequest:\n%+v", err, string(reqJson))
		return apiResponse, err
	}

	return apiResponse, nil
}

/***********************************************************************************************************************
 * PUT /api/v1/teams/roles/:id
 **********************************************************************************************************************/

type UpdateTeamRoleRequest struct {
	ID                   int64  `json:"id"`
	Tag                  string `json:"tag"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	NotificationsEnabled bool   `json:"notificationsEnabled,omitempty"`
}

// ToUpdateRequest converts a TeamRole to an UpdateTeamRoleRequest.
func (r *TeamRole) ToUpdateRequest() UpdateTeamRoleRequest {
	return UpdateTeamRoleRequest{
		ID:                   r.ID,
		Tag:                  r.Tag,
		Name:                 r.Name,
		Description:          r.Description,
		NotificationsEnabled: r.NotificationsEnabled,
	}
}

func (c *TeamRolesClient) Update(_ context.Context, req UpdateTeamRoleRequest) (*TeamRole, error) {
	apiResponse := &TeamRole{}
	apiError := &ApiError{}

	response, err := c.Client().Put(Route("teams_roles", fmt.Sprintf("%d", req.ID))).BodyJSON(&req).Receive(apiResponse, apiError)
	if err != nil {
		return apiResponse, errors.New("could not update team role: " + err.Error())
	}

	err = c.client.handleResponseStatus(response, apiError)
	if err != nil {
		reqJson, _ := json.Marshal(req)
		log.Printf("Failed updating team role: %+v\n\nRequest:\n%+v\n%+v", err, string(reqJson), apiError.String())
		return apiResponse, err
	}

	return apiResponse, nil
}

/***********************************************************************************************************************
 * DELETE /api/v1/teams/roles/:id
 **********************************************************************************************************************/

type DeleteTeamRoleResponse struct{}
type DeleteTeamRoleRequest struct {
	ID int64 `json:"id"`
}

func (c *TeamRolesClient) Delete(_ context.Context, id int64) error {
	apiError := &ApiError{}
	apiResponse := &DeleteTeamRoleResponse{}

	req := DeleteTeamRoleRequest{ID: id}
	resp, err := c.Client().Delete(Route("teams_roles", fmt.Sprintf("%d", id))).QueryStruct(req).Receive(apiResponse, apiError)
	if err != nil {
		return fmt.Errorf("could not delete team role %d:\n\n%+v", id, err.Error())
	}

	err = c.client.handleResponseStatus(resp, apiError)
	if err != nil {
		log.Printf("Could not delete team role %d:\n\n%+v", id, err.Error())
		return err
	}

	return nil
}
