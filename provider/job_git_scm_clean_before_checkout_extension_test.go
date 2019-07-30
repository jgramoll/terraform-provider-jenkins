package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitScmCleanBeforeCheckoutExtensionBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	// extensionResourceName := "jenkins_job_git_scm_clean_before_checkout_extension.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
resource "jenkins_job" "main" {
	name = "%s"
}

resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.id}"
}

resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
  job = "${jenkins_job.main.id}"
  scm = "${jenkins_job_git_scm.main.id}"
}`, jobName)
}
