package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitScmCleanBeforeCheckoutExtensionBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.test"
	// extensionResourceName := "jenkins_job_git_scm_clean_before_checkout_extension.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccJobGitScmCleanBeforeCheckoutExtensionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "test" {
	name = "%s"
}

resource "jenkins_job_git_scm" "test" {
	job = "${jenkins_job.test.id}"
}

resource "jenkins_job_git_scm_clean_before_checkout_extension" "test" {
  job = "${jenkins_job.test.id}"
  scm = "${jenkins_job_git_scm.test.id}"
}`, jobName)
}

func testAccJobGitScmCleanBeforeCheckoutExtensionDestroy(s *terraform.State) error {
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
