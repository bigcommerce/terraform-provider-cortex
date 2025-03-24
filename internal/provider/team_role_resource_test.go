package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamRoleMinimalResource(t *testing.T) {
	tag := "test-team-role-minimal"
	resourceType := "cortex_team_role"
	resourceName := resourceType + "." + tag
	stub := tFactoryBuildTeamRoleMinimalResource(tag)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTeamRoleMinimalResourceConfig(resourceType, stub),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tag", stub.Tag),
					resource.TestCheckResourceAttr(resourceName, "name", stub.Name),
					resource.TestCheckResourceAttr(resourceName, "description", stub.Description),
					resource.TestCheckResourceAttr(resourceName, "notifications_enabled", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccTeamRoleMinimalResourceConfig(resourceType, stub),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tag", stub.Tag),
					resource.TestCheckResourceAttr(resourceName, "name", stub.Name),
					resource.TestCheckResourceAttr(resourceName, "description", stub.Description),
					resource.TestCheckResourceAttr(resourceName, "notifications_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccTeamRoleMinimalResourceConfig(resourceType string, stub TestTeamRoleResource) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
 	tag = %[3]q
 	name = %[4]q
 	description = %[5]q
    notifications_enabled = true
}
`, resourceType, stub.Tag, stub.Tag, stub.Name, stub.Description)
}

type TestTeamRoleResource struct {
	Tag                  string
	Name                 string
	Description          string
	NotificationsEnabled bool
}

func tFactoryBuildTeamRoleMinimalResource(tag string) TestTeamRoleResource {
	return TestTeamRoleResource{
		Tag:                  tag,
		Name:                 "Test Role",
		Description:          "A test role",
		NotificationsEnabled: true,
	}
}
