package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritProjectBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	pattern := "project-1"
	newPattern := "new-project"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritProjectConfigPattern(jobName, pattern),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.compare_type", "REG_EXP"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.pattern", pattern),
				),
			},
			{
				Config: testAccJobGerritProjectConfigPattern(jobName, newPattern),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.compare_type", "REG_EXP"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.gerrit_project.0.pattern", newPattern),
				),
			},
		},
	})
}

func testAccJobGerritProjectConfigPattern(jobName string, projectPattern string) string {
	return fmt.Sprintf(`
	resource "jenkins_job" "main" {
		name     = "%s"
		plugin   = "workflow-job@2.33"
	
		property {
			type = "PipelineTriggersJobProperty"
	
			trigger {
				type   = "GerritTrigger"
				skip_vote {}

				gerrit_project {
					compare_type = "REG_EXP"
					pattern      = "%s"
				}
			}
		}
	}`, jobName, projectPattern)
}
