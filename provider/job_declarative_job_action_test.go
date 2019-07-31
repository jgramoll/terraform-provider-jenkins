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

func TestAccJobDeclarativeJobActionBasic(t *testing.T) {
	var jobRef client.Job
	var actions []client.JobAction
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	actionResourceName := "jenkins_job_declarative_job_action.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDeclarativeJobActionConfig(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobActions(&jobRef, []string{
						actionResourceName,
					}, &actions, ensureJobDeclarativeJobAction),
				),
			},
			{
				Config: testAccJobConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobActions(&jobRef, []string{}, &actions, ensureJobDeclarativeJobAction),
				),
			},
		},
	})
}

func testAccJobDeclarativeJobActionConfig(jobName string) string {
	return testAccJobConfigBasic(jobName) + `
resource "jenkins_job_declarative_job_action" "main" {
  job = "${jenkins_job.main.id}"
}
`
}

func ensureJobDeclarativeJobAction(actionInterface client.JobAction, rs *terraform.ResourceState) error {
	_, ok := actionInterface.(*client.JobDeclarativeJobAction)
	if !ok {
		return fmt.Errorf("Action is not of expected type, expected *client.JobDeclarativeJobAction, actually %s",
			reflect.TypeOf(actionInterface).String())
	}

	return nil
}
