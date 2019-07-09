package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobBasic(t *testing.T) {
	var jobRef client.Job
	folder := "job/Bridge%20Career"
	name := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	resourceName := "jenkins_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobConfigBasic(folder, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(resourceName, &jobRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "folder", folder),
				),
			},
			{
				Config: testAccJobConfigBasic(folder, newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(resourceName, &jobRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
					resource.TestCheckResourceAttr(resourceName, "folder", folder),
				),
			},
		},
	})
}

func testAccJobConfigBasic(folder string, name string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "test" {
  folder = "%s"
  name   = "%s"
}`, folder, name)
}

func testAccCheckJobExists(resourceName string, j *client.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		jobService := testAccProvider.Meta().(*Services).JobService
		job, err := jobService.GetJob(rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}
		*j = *job

		return nil
	}
}

func testAccCheckJobDestroy(s *terraform.State) error {
	jobService := testAccProvider.Meta().(*Services).JobService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "jenkins_job" {
			_, err := jobService.GetJob(rs.Primary.Attributes["name"])
			if err == nil {
				return fmt.Errorf("Job still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
