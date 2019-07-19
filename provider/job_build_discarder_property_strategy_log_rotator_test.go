package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobBuildDiscarderPropertyStrategyTypes["jenkins_job_build_discarder_property_log_rotator_strategy"] = reflect.TypeOf((*client.JobBuildDiscarderPropertyStrategyLogRotator)(nil))
}

func TestAccJobBuildDiscarderPropertyStrategyLogRotatorBasic(t *testing.T) {
	var jobRef client.Job
	var strategyRefs []client.JobBuildDiscarderPropertyStrategy
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	strategyResourceName := "jenkins_job_build_discarder_property_log_rotator_strategy.main"
	daysToKeep := "1"
	newDaysToKeep := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, daysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(strategyResourceName, "days_to_keep", daysToKeep),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckBuildDiscarderPropertyStrategies(&jobRef, []string{
						strategyResourceName,
					}, &strategyRefs),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, newDaysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(strategyResourceName, "days_to_keep", newDaysToKeep),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckBuildDiscarderPropertyStrategies(&jobRef, []string{
						strategyResourceName,
					}, &strategyRefs),
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
