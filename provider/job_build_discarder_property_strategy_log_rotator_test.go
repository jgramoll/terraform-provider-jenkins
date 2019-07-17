package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobBuildDiscarderPropertyStrategyLogRotatorBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.test"
	strategyResourceName := "jenkins_job_build_discarder_property_log_rotator_strategy.test"
	daysToKeep := "1"
	newDaysToKeep := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccJobBuildDiscarderPropertyStrategyLogRotatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, daysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(strategyResourceName, "days_to_keep", daysToKeep),
					// TODO test daysToKeep on job
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, newDaysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(strategyResourceName, "days_to_keep", newDaysToKeep),
					// testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName string, daysToKeep string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name = "%s"
}

resource "jenkins_job_build_discarder_property" "main" {
  job = "${jenkins_job.main.id}"
}

resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
  property = "${jenkins_job_build_discarder_property.main.id}"

  days_to_keep          = %s
  num_to_keep           = -1
  artifact_days_to_keep = -1
  artifact_num_to_keep  = -1
}`, jobName, daysToKeep)
}

func testAccJobBuildDiscarderPropertyStrategyLogRotatorDestroy(s *terraform.State) error {
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
