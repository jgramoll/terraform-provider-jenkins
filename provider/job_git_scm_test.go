package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

// func init() {
// 	stageTypes["spinnaker_pipeline_destroy_server_group_stage"] = client.DestroyServerGroupStageType
// }

func TestAccJobGitScmBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	scriptPath := "my-script"
	newScriptPath := "new-my-script"
	jobResourceName := "jenkins_job.test"
	definition := "jenkins_job_git_scm.test"
	expected := client.CpsScmFlowDefinition{}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmConfigBasic(jobName, scriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "script_path", scriptPath),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobDefinition(jobResourceName, &expected, &jobRef),
				),
			},
			{
				Config: testAccJobGitScmConfigBasic(jobName, newScriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "script_path", newScriptPath),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobDefinition(jobResourceName, &expected, &jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmConfigBasic(jobName string, scriptPath string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "test" {
	name = "%s"
}

resource "jenkins_job_git_scm" "test" {
	job = "${jenkins_job.test.id}"

	script_path = "%s"
}`, jobName, scriptPath)
}
