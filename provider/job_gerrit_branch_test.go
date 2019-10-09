package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritBranchBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	compareType := "PLAIN"
	newCompareType := "REG_EXP"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritBranchConfigBasic(jobName, compareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.branch.0.compare_type", compareType),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.branch.0.pattern", "branch_pattern"),
				),
			},
			{
				Config: testAccJobGerritBranchConfigBasic(jobName, newCompareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.branch.0.compare_type", newCompareType),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.branch.0.pattern", "branch_pattern"),
				),
			},
		},
	})
}

func testAccJobGerritBranchConfigBasic(jobName string, compareType string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

	property {
		type = "PipelineTriggersJobProperty"

		trigger {
			type   = "GerritTrigger"
			plugin = "gerrit-trigger@2.29.0"
			skip_vote {}

			gerrit_project {
				compare_type = "REG_EXP"
				pattern      = "gerrit_project"

				branch {
					compare_type = "%s"
					pattern      = "branch_pattern"
				}
			}
		}
	}
}`, jobName, compareType)
}
