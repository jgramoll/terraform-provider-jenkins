package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

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
					}, &events, ensureJobGerritTriggerDraftPublishedEvent),
				),
			},
			{
				ResourceName:  eventResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid trigger event id"),
			},
			{
				ResourceName: eventResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(events) == 0 {
						return "", fmt.Errorf("no gerrit trigger event to import")
					}
					property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
					trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)
					return strings.Join([]string{jobName, property.Id, trigger.Id, events[0].GetId()}, IdDelimiter), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobGerritTriggerConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritTriggerEvents(&jobRef, []string{}, &events, ensureJobGerritTriggerDraftPublishedEvent),
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

func ensureJobGerritTriggerDraftPublishedEvent(
	eventInterface client.JobGerritTriggerOnEvent,
	rs *terraform.ResourceState,
) error {
	_, err := newJobGerritTriggerDraftPublishedEvent().fromClientJobTriggerEvent(eventInterface)
	return err
}
