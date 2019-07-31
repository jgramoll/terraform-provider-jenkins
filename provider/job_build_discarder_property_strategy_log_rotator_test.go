package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobBuildDiscarderPropertyStrategyTypes["jenkins_job_build_discarder_property_log_rotator_strategy"] = reflect.TypeOf((*client.JobBuildDiscarderPropertyStrategyLogRotator)(nil))
}

func TestAccJobBuildDiscarderPropertyStrategyLogRotatorBasic(t *testing.T) {
	var jobRef client.Job
	var strategyRefs []client.JobBuildDiscarderPropertyStrategy
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
					}, &strategyRefs, ensureJobBuildDiscarderPropertyStrategyLogRotator),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, newDaysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(strategyResourceName, "days_to_keep", newDaysToKeep),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckBuildDiscarderPropertyStrategies(&jobRef, []string{
						strategyResourceName,
					}, &strategyRefs, ensureJobBuildDiscarderPropertyStrategyLogRotator),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckBuildDiscarderPropertyStrategies(&jobRef, []string{}, &strategyRefs, ensureJobBuildDiscarderPropertyStrategyLogRotator),
				),
			},
		},
	})
}

func testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName string, daysToKeep string) string {
	return testAccJobBuildDiscarderPropertyConfigBasic(jobName, 1) +
		fmt.Sprintf(`
resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
  property = "${jenkins_job_build_discarder_property.prop_1.id}"

  days_to_keep          = %s
  num_to_keep           = -1
  artifact_days_to_keep = -1
  artifact_num_to_keep  = -1
}`, daysToKeep)
}

func ensureJobBuildDiscarderPropertyStrategyLogRotator(strategyInterface client.JobBuildDiscarderPropertyStrategy, rs *terraform.ResourceState) error {
	strategy, ok := strategyInterface.(*client.JobBuildDiscarderPropertyStrategyLogRotator)
	if !ok {
		return fmt.Errorf("Strategy is not of expected type, expected *client.JobBuildDiscarderPropertyStrategyLogRotator, actually %s",
			reflect.TypeOf(strategyInterface).String())
	}

	_, _, strategyId, err := resourceJobPropertyStrategyId(rs.Primary.Attributes["id"])
	if err != nil {
		return err
	}
	if strategyId != strategy.Id {
		return fmt.Errorf("JobBuildDiscarderPropertyStrategyLogRotator id should be %v, was %v", strategyId, strategy.Id)
	}
	err = testCompareResourceInt("JobBuildDiscarderPropertyStrategyLogRotator", "DaysToKeep", rs.Primary.Attributes["days_to_keep"], strategy.DaysToKeep)
	if err != nil {
		return err
	}
	err = testCompareResourceInt("JobBuildDiscarderPropertyStrategyLogRotator", "NumToKeep", rs.Primary.Attributes["num_to_keep"], strategy.NumToKeep)
	if err != nil {
		return err
	}
	err = testCompareResourceInt("JobBuildDiscarderPropertyStrategyLogRotator", "ArtifactDaysToKeep", rs.Primary.Attributes["artifact_days_to_keep"], strategy.ArtifactDaysToKeep)
	if err != nil {
		return err
	}
	err = testCompareResourceInt("JobBuildDiscarderPropertyStrategyLogRotator", "ArtifactNumToKeep", rs.Primary.Attributes["artifact_num_to_keep"], strategy.ArtifactNumToKeep)
	if err != nil {
		return err
	}

	return nil
}
