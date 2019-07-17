package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobBuildDiscarderPropertyBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	// propertyResourceName := "jenkins_job_build_discarder_property.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccJobBuildDiscarderPropertyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobBuildDiscarderPropertyConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					// TODO check that property exists
				),
			},
		},
	})
}

func testAccJobBuildDiscarderPropertyConfigBasic(jobName string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name = "%s"
}

resource "jenkins_job_build_discarder_property" "main" {
  job = "${jenkins_job.main.id}"
}`, jobName)
}

func testAccJobBuildDiscarderPropertyDestroy(s *terraform.State) error {
	jobService := testAccProvider.Meta().(*Services).JobService
	for _, rs := range s.RootModule().Resources {
		if _, ok := jobPropertyTypes[rs.Type]; ok {
			_, err := jobService.GetJob(rs.Primary.Attributes["name"])
			// TODO does this really check anything?
			if err == nil {
				return fmt.Errorf("Job Git Scm User Remote Config still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
