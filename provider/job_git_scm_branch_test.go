package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitBranchBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	branchName := "my-branch"
	newBranchName := "new-my-branch"
	jobResourceName := "jenkins_job.main"
	definition := "jenkins_job_git_scm_branch.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobBranchDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmBranchConfigBasic(jobName, branchName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "name", branchName),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
			{
				Config: testAccJobGitScmBranchConfigBasic(jobName, newBranchName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "name", newBranchName),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmBranchConfigBasic(jobName string, branchName string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name = "%s"
}

resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"
}


resource "jenkins_job_git_scm_branch" "main" {
	job = "${jenkins_job.main.id}"
  scm = "${jenkins_job_git_scm.main.id}"

  name = "%s"
}`, jobName, branchName)
}

func testAccCheckJobBranchDestroy(s *terraform.State) error {
	jobService := testAccProvider.Meta().(*Services).JobService
	for _, rs := range s.RootModule().Resources {
		if _, ok := jobPropertyTypes[rs.Type]; ok {
			_, err := jobService.GetJob(rs.Primary.Attributes["name"])
			// TODO does this really check anything?
			if err == nil {
				return fmt.Errorf("Job Git Scm User Remote Config still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
