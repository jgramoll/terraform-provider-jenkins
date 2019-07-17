package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritTriggerPatchsetCreatedEventBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.test"
	propertyResourceName := "jenkins_job_build_discarder_property.test"
	strategy := "hudson.tasks.LogRotator"
	newStrategy := "hudson.tasks.LogRotatorNew"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccJobGerritTriggerPatchsetCreatedEventDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName, strategy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(propertyResourceName, "strategy", strategy),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
			{
				Config: testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName, newStrategy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(propertyResourceName, "strategy", newStrategy),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName string, strategy string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "test" {
	name = "%s"
}

resource "jenkins_job_build_discarder_property" "test" {
  job = "${jenkins_job.test.id}"

  strategy              = "%s"
  days_to_keep          = "1"
  num_to_keep           = "2"
  artifact_days_to_keep = "3"
  artifact_num_to_keep  = "4"
}`, jobName, strategy)
}

func testAccJobGerritTriggerPatchsetCreatedEventDestroy(s *terraform.State) error {
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
