package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitScmUserRemoteConfigBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	refspec := "my-refspec"
	newRefspec := "new-my-refspec"
	jobResourceName := "jenkins_job.test"
	definition := "jenkins_job_git_scm_user_remote_config.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobUserRemoteConfigDestroy,
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
resource "jenkins_job" "test" {
	name = "%s"
}

resource "jenkins_job_git_scm" "test" {
	job = "${jenkins_job.test.id}"
}

resource "jenkins_job_git_scm_user_remote_config" "test" {
	job = "${jenkins_job.test.id}"
	scm = "${jenkins_job_git_scm.test.id}"

  refspec        = "%s"
  url            = "my-test-url"
  credentials_id = "my-test-creds"
}`, jobName, refspec)
}

func testAccCheckJobUserRemoteConfigDestroy(s *terraform.State) error {
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
