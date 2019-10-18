package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritTriggerBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, "gerrit-trigger@2.28.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.type", "GerritTrigger"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.plugin", "gerrit-trigger@2.28.0"),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, "gerrit-trigger@2.29.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.type", "GerritTrigger"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.plugin", "gerrit-trigger@2.29.0"),
				),
			},
		},
	})
}

func testAccJobGerritTriggerConfigServerName(jobName string, plugin string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

	property {
		type = "PipelineTriggersJobProperty"

		trigger {
			type   = "GerritTrigger"
			plugin = "%s"
			skip_vote {}
		}
	}
}`, jobName, plugin)
}
