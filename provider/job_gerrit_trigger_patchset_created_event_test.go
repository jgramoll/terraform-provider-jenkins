package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritTriggerPatchsetCreatedEventBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.type", "PluginPatchsetCreatedEvent"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.exclude_drafts", "false"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.exclude_trivial_rebase", "false"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.exclude_no_code_change", "false"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.exclude_private_state", "false"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.trigger.0.trigger_on_event.0.exclude_wip_state", "false"),
				),
			},
		},
	})
}

func testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName string, excludeDrafts string) string {
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

			trigger_on_event {
				type = "PluginPatchsetCreatedEvent"

				exclude_drafts         = false
				exclude_trivial_rebase = false
				exclude_no_code_change = false
				exclude_private_state  = false
				exclude_wip_state      = false
			}
		}
	}
}`, jobName)
}
