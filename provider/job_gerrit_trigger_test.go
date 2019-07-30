package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerTypes["jenkins_job_gerrit_trigger"] = reflect.TypeOf((*client.JobGerritTrigger)(nil))
}

func TestAccJobGerritTriggerBasic(t *testing.T) {
	var jobRef client.Job
	var triggers []client.JobTrigger
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	trigger1 := "jenkins_job_gerrit_trigger.trigger_1"
	trigger2 := "jenkins_job_gerrit_trigger.trigger_2"
	serverName := "my-server"
	newServerName := "my-new-server"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					resource.TestCheckResourceAttr(trigger2, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger2, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, newServerName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "server_name", newServerName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					resource.TestCheckResourceAttr(trigger2, "server_name", newServerName),
					resource.TestCheckResourceAttr(trigger2, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "server_name", serverName),
					resource.TestCheckResourceAttr(trigger1, "silent_mode", "false"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigServerName(jobName, serverName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobTriggers(&jobRef, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccJobGerritTriggerConfigBasic(jobName string) string {
	return testAccJobGerritTriggerConfigServerName(jobName, "__ANY__", 1)
}

func testAccJobGerritTriggerConfigServerName(jobName string, serverName string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger" "trigger_%v" {
	property = "${jenkins_job_pipeline_triggers_property.prop_1.id}"

	server_name = "%v"

	skip_vote {}
}`, i, serverName)
	}

	t := testAccJobPipelineTriggersPropertyConfigBasic(jobName, 1) + triggers
	return t
}
