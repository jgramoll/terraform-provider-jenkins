package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobBuildDiscarderPropertyStrategyLogRotatorBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	daysToKeep := "1"
	newDaysToKeep := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, daysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.type", "LogRotator"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.days_to_keep", daysToKeep),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.num_to_keep", "-1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.artifact_days_to_keep", "-1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.artifact_num_to_keep", "-1"),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName, newDaysToKeep),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.type", "LogRotator"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.days_to_keep", newDaysToKeep),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.num_to_keep", "-1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.artifact_days_to_keep", "-1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.strategy.0.artifact_num_to_keep", "-1"),
				),
			},
		},
	})
}

func testAccJobBuildDiscarderPropertyStrategyLogRotatorConfigBasic(jobName string, daysToKeep string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

	property {
		type = "BuildDiscarderProperty"

		strategy {
			type  = "LogRotator"

			days_to_keep          = "%s"
			num_to_keep           = "-1"
			artifact_days_to_keep = "-1"
			artifact_num_to_keep  = "-1"
		}
	}
}`, jobName, daysToKeep)
}
