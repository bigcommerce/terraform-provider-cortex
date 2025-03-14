package provider

import (
	"github.com/cortexapps/terraform-provider-cortex/internal/cortex"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/***********************************************************************************************************************
 * Models
 **********************************************************************************************************************/

// TeamRoleResourceModel describes the team role data model within Terraform.
type TeamRoleResourceModel struct {
	ID                   types.Int64  `tfsdk:"id"`
	Tag                  types.String `tfsdk:"tag"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	NotificationsEnabled types.Bool   `tfsdk:"notifications_enabled"`
}

func (r *TeamRoleResourceModel) FromApiModel(entity cortex.TeamRole) {
	r.ID = types.Int64Value(entity.ID)
	r.Tag = types.StringValue(entity.Tag)
	r.Name = types.StringValue(entity.Name)
	if entity.Description != "" {
		r.Description = types.StringValue(entity.Description)
	} else {
		r.Description = types.StringNull()
	}
	if entity.NotificationsEnabled {
		r.NotificationsEnabled = types.BoolValue(true)
	} else {
		r.NotificationsEnabled = types.BoolValue(false)
	}
}

func (r *TeamRoleResourceModel) ToApiModel() cortex.TeamRole {
	entity := cortex.TeamRole{
		ID:                   r.ID.ValueInt64(),
		Tag:                  r.Tag.ValueString(),
		Name:                 r.Name.ValueString(),
		Description:          r.Description.ValueString(),
		NotificationsEnabled: r.NotificationsEnabled.ValueBool(),
	}
	return entity
}

/***********************************************************************************************************************
 * Data Source
 **********************************************************************************************************************/

// TeamRoleDataSourceModel describes the data source data model.
type TeamRoleDataSourceModel struct {
	ID                   types.Int64  `tfsdk:"id"`
	Tag                  types.String `tfsdk:"tag"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	NotificationsEnabled types.Bool   `tfsdk:"notifications_enabled"`
}

func (o *TeamRoleDataSourceModel) FromApiModel(entity cortex.TeamRole) {
	o.ID = types.Int64Value(entity.ID)
	o.Tag = types.StringValue(entity.Tag)
	o.Name = types.StringValue(entity.Name)
	o.Description = types.StringValue(entity.Description)
	o.NotificationsEnabled = types.BoolValue(entity.NotificationsEnabled)
}
