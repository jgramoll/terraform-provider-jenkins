package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritFilePathBasic(t *testing.T) {
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
				Config: testAccJobGerritFilePathConfigBasic(jobName, compareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.file_path.0.compare_type", compareType),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.file_path.0.pattern", "file_path"),
				),
			},
			{
				Config: testAccJobGerritFilePathConfigBasic(jobName, newCompareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.file_path.0.compare_type", newCompareType),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.file_path.0.pattern", "file_path"),
				),
			},
		},
	})
}

func testAccJobGerritFilePathConfigBasic(jobName string, compareType string) string {
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

				file_path {
					compare_type = "%s"
					pattern      = "file_path"
				}
			}
		}
	}
}`, jobName, compareType)
}
