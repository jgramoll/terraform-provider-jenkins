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
	jobTriggerEventTypes["jenkins_job_gerrit_trigger_patchset_created_event"] = reflect.TypeOf((*client.JobGerritTriggerPluginPatchsetCreatedEvent)(nil))
}

func TestAccJobGerritTriggerPatchsetCreatedEventBasic(t *testing.T) {
	var jobRef client.Job
	var events []client.JobGerritTriggerOnEvent
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	eventResourceName := "jenkins_job_gerrit_trigger_patchset_created_event.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(eventResourceName, "exclude_drafts", "true"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{
						eventResourceName,
					}, &events, ensureJobGerritTriggerPatchsetCreatedEvent),
				),
			},
			{
				Config: testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName, "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(eventResourceName, "exclude_drafts", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{
						eventResourceName,
					}, &events, ensureJobGerritTriggerPatchsetCreatedEvent),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{}, &events, ensureJobGerritTriggerPatchsetCreatedEvent),
				),
			},
		},
	})
}

func testAccJobGerritTriggerPatchsetCreatedEventConfigBasic(jobName string, excludeDrafts string) string {
	return testAccJobGerritTriggerConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1.id}"
	
	exclude_drafts = %s
}`, excludeDrafts)
}

func ensureJobGerritTriggerPatchsetCreatedEvent(
	eventInterface client.JobGerritTriggerOnEvent,
	rs *terraform.ResourceState,
) error {
	event, ok := eventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	if !ok {
		return fmt.Errorf("Unexpected event type got %s, expected *client.JobGerritTriggerPluginPatchsetCreatedEvent",
			reflect.TypeOf(eventInterface).String())
	}

	err := testCompareResourceBool("JobGerritTriggerPluginPatchsetCreatedEvent", "ExcludeDrafts", rs.Primary.Attributes["exclude_drafts"], event.ExcludeDrafts)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTriggerPluginPatchsetCreatedEvent", "ExcludeTrivialRebase", rs.Primary.Attributes["exclude_trivial_rebase"], event.ExcludeTrivialRebase)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTriggerPluginPatchsetCreatedEvent", "ExcludeNoCodeChange", rs.Primary.Attributes["exclude_no_code_change"], event.ExcludeNoCodeChange)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTriggerPluginPatchsetCreatedEvent", "ExcludePrivateState", rs.Primary.Attributes["exclude_private_state"], event.ExcludePrivateState)
	if err != nil {
		return err
	}
	err = testCompareResourceBool("JobGerritTriggerPluginPatchsetCreatedEvent", "ExcludeWipState", rs.Primary.Attributes["exclude_wip_state"], event.ExcludeWipState)
	if err != nil {
		return err
	}

	return nil
}
