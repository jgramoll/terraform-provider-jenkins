package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitScmBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	definitionResourceName := "jenkins_job_git_scm.main"
	scriptPath := "my-script"
	newScriptPath := "new-my-script"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmConfigScript(jobName, scriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definitionResourceName, "script_path", scriptPath),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobDefinition(&jobRef, definitionResourceName, testAccJobGitScmEnsureDefinition),
				),
			},
			{
				Config: testAccJobGitScmConfigScript(jobName, newScriptPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definitionResourceName, "script_path", newScriptPath),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobDefinition(&jobRef, definitionResourceName, testAccJobGitScmEnsureDefinition),
				),
			},
			{
				Config: testAccJobConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckNoJobDefinition(&jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmConfigBasic(jobName string) string {
	return testAccJobConfigBasic(jobName) + `
resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"
}`
}

func testAccJobGitScmConfigScript(jobName string, scriptPath string) string {
	return testAccJobConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"

	script_path = "%s"
}`, scriptPath)
}

func testAccJobGitScmEnsureDefinition(definitionInterface client.JobDefinition, rs *terraform.ResourceState) error {
	definition := definitionInterface.(*client.CpsScmFlowDefinition)

	_, definitionId, err := resourceJobDefinitionId(rs.Primary.Attributes["id"])
	if err != nil {
		return err
	}
	if definitionId != definition.Id {
		return fmt.Errorf("CpsScmFlowDefinition id should be %v, was %v", definitionId, definition.Id)
	}
	if rs.Primary.Attributes["script_path"] != definition.ScriptPath {
		return fmt.Errorf("CpsScmFlowDefinition script_path should be %v, was %v", rs.Primary.Attributes["script_path"], definition.ScriptPath)
	}
	err = testCompareResourceBool("CpsScmFlowDefinition", "lightweight", rs.Primary.Attributes["lightweight"], definition.Lightweight)
	if err != nil {
		return err
	}
	if rs.Primary.Attributes["config_version"] != definition.SCM.ConfigVersion {
		return fmt.Errorf("CpsScmFlowDefinition config_version should be %v, was %v", rs.Primary.Attributes["config_version"], definition.SCM.ConfigVersion)
	}

	return nil
}
