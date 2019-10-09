package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobCpsCsmFlowDefinitionBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	scriptPath := "my-script"
	newScriptPath := "new-my-script"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobCpsCsmFlowDefinitionConfigBasic(jobName, scriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.type", "CpsScmFlowDefinition"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.plugin", "workflow-cps@2.70"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.script_path", scriptPath),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.lightweight", "false"),
				),
			},
			{
				Config: testAccJobCpsCsmFlowDefinitionConfigBasic(jobName, newScriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.type", "CpsScmFlowDefinition"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.plugin", "workflow-cps@2.70"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.script_path", newScriptPath),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.lightweight", "false"),
				),
			},
		},
	})
}

func testAccJobCpsCsmFlowDefinitionConfigBasic(jobName string, scriptPath string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name   = "%s"
	plugin = "workflow-job@2.33"

	definition {
		type   = "CpsScmFlowDefinition"
		plugin = "workflow-cps@2.70"

		script_path = "%s"
		lightweight  = false
	}
}`, jobName, scriptPath)
}
