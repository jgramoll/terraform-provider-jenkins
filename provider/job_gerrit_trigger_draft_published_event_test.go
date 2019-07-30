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
	jobTriggerEventTypes["jenkins_job_gerrit_trigger_draft_published_event"] = reflect.TypeOf((*client.JobGerritTriggerPluginDraftPublishedEvent)(nil))
}

func TestAccJobGerritTriggerDraftPublishedEventBasic(t *testing.T) {
	var jobRef client.Job
	var events []client.JobGerritTriggerOnEvent
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	eventResourceName := "jenkins_job_gerrit_trigger_draft_published_event.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerDraftPublishedEventConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{
						eventResourceName,
					}, &events, testAccEnsureJobGerritTriggerDraftPublishedEvent),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{}, &events, testAccEnsureJobGerritTriggerDraftPublishedEvent),
				),
			},
		},
	})
}

func testAccJobGerritTriggerDraftPublishedEventConfigBasic(jobName string) string {
	return testAccJobGerritTriggerConfigBasic(jobName) + `
resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
  trigger = "${jenkins_job_gerrit_trigger.trigger_1.id}"
}`
}

func testAccEnsureJobGerritTriggerDraftPublishedEvent(
	eventInterface client.JobGerritTriggerOnEvent,
	rs *terraform.ResourceState,
) error {
	event := eventInterface.(*client.JobGerritTriggerPluginDraftPublishedEvent)

	_, _, _, eventId, err := resourceJobTriggerEventId(rs.Primary.Attributes["id"])
	if err != nil {
		return err
	}
	if eventId != event.Id {
		return fmt.Errorf("testAccEnsureJobGerritTriggerDraftPublishedEvent id should be %v, was %v", eventId, event.Id)
	}

	return nil
}
