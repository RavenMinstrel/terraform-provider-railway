package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccBucketResourceDefault(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccBucketResourceConfigDefault("integration"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("railway_bucket.test", "id", uuidRegex()),
					resource.TestCheckResourceAttr("railway_bucket.test", "name", "integration"),
					resource.TestCheckResourceAttr("railway_bucket.test", "project_id", "0bb01547-570d-4109-a5e8-138691f6a2d1"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "railway_bucket.test",
				ImportState:       true,
				ImportStateIdFunc: testAccBucketImportStateIdFunc("railway_bucket.test"),
				ImportStateVerify: true,
			},
			// Update with new name
			{
				Config: testAccBucketResourceConfigDefault("integration-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("railway_bucket.test", "id", uuidRegex()),
					resource.TestCheckResourceAttr("railway_bucket.test", "name", "integration-updated"),
					resource.TestCheckResourceAttr("railway_bucket.test", "project_id", "0bb01547-570d-4109-a5e8-138691f6a2d1"),
				),
			},
			// Delete testing automatically occurs in TestCase
			// NOTE: Railway does not support deleting buckets via API.
			// The bucket will be removed from state but will still exist in Railway.
		},
	})
}

func testAccBucketResourceConfigDefault(name string) string {
	return fmt.Sprintf(`
resource "railway_bucket" "test" {
  name       = "%s"
  project_id = "0bb01547-570d-4109-a5e8-138691f6a2d1"
}
`, name)
}

func testAccBucketImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return rs.Primary.Attributes["project_id"] + ":" + rs.Primary.ID, nil
	}
}
