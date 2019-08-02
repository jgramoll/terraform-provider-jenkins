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

func TestAccJobDeclarativeJobPropertyTrackerActionBasic(t *testing.T) {
	var jobRef client.Job
	var actions []client.JobAction
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	actionResourceName := "jenkins_job_declarative_job_property_tracker_action.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDeclarativeJobPropertyTrackerActionConfig(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobActions(&jobRef, []string{
						actionResourceName,
					}, &actions, ensureJobDeclarativeJobPropertyTrackerAction),
				),
			},
			{
				ResourceName:  actionResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid action id"),
			},
			{
				ResourceName: actionResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(actions) == 0 {
						return "", fmt.Errorf("no actions to import")
					}
					return strings.Join([]string{jobName, actions[0].GetId()}, IdDelimiter), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobActions(&jobRef, []string{}, &actions, ensureJobDeclarativeJobPropertyTrackerAction),
				),
			},
		},
	})
}

func testAccJobDeclarativeJobPropertyTrackerActionConfig(jobName string) string {
	return testAccJobConfigBasic(jobName) + `
resource "jenkins_job_declarative_job_property_tracker_action" "main" {
  job = "${jenkins_job.main.id}"
}
`
}

func ensureJobDeclarativeJobPropertyTrackerAction(actionInterface client.JobAction, rs *terraform.ResourceState) error {
	_, err := newJobDeclarativeJobPropertyTrackerAction().fromClientAction(actionInterface)
	return err
}
