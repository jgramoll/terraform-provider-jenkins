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
	name := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	resourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(resourceName, &jobRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccJobConfigBasic(newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(resourceName, &jobRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
				),
			},
		},
	})
}

func testAccJobConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
  name   = "%s"
}`, name)
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
