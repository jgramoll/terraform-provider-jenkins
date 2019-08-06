package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

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
				ResourceName:  strategyResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid discarder property strategy id"),
			},
			{
				ResourceName: strategyResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(strategyRefs) == 0 {
						return "", fmt.Errorf("no strategies to import")
					}
					propertyId := (*jobRef.Properties.Items)[0].GetId()
					return ResourceJobDiscarderPropertyStrategyId(jobName, propertyId, strategyRefs[0].GetId()), nil
				},
				ImportStateVerify: true,
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

func ensureJobBuildDiscarderPropertyStrategyLogRotator(clientStrategyInterface client.JobBuildDiscarderPropertyStrategy, rs *terraform.ResourceState) error {
	strategyInterface, err := newJobBuildDiscarderPropertyStrategyLogRotator().fromClientStrategy(clientStrategyInterface)
	if err != nil {
		return err
	}
	strategy := strategyInterface.(*jobBuildDiscarderPropertyStrategyLogRotator)

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
