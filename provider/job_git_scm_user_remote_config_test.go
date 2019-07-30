package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitScmUserRemoteConfigBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	refspec := "my-refspec"
	newRefspec := "new-my-refspec"
	jobResourceName := "jenkins_job.main"
	definition := "jenkins_job_git_scm_user_remote_config.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, refspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "refspec", refspec),
					resource.TestCheckResourceAttr(definition, "url", "my-test-url"),
					resource.TestCheckResourceAttr(definition, "credentials_id", "my-test-creds"),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, newRefspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(definition, "refspec", newRefspec),
					resource.TestCheckResourceAttr(definition, "url", "my-test-url"),
					resource.TestCheckResourceAttr(definition, "credentials_id", "my-test-creds"),
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmUserRemoteConfigConfigBasic(jobName string, refspec string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name = "%s"
}

resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_git_scm_user_remote_config" "main" {
	job = "${jenkins_job.main.id}"
	scm = "${jenkins_job_git_scm.main.id}"

  refspec        = "%s"
  url            = "my-test-url"
  credentials_id = "my-test-creds"
}`, jobName, refspec)
}
