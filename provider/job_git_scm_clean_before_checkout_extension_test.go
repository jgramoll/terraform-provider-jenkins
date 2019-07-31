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

func init() {
	jobGitScmExtensionTypes["jenkins_job_git_scm_clean_before_checkout_extension"] = reflect.TypeOf((*client.GitScmCleanBeforeCheckoutExtension)(nil))
}

func TestAccJobGitScmCleanBeforeCheckoutExtensionBasic(t *testing.T) {
	var jobRef client.Job
	var extensions []client.GitScmExtension
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	extensionResourceName := "jenkins_job_git_scm_clean_before_checkout_extension.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmExtensions(&jobRef, []string{
						extensionResourceName,
					}, &extensions, ensureJobGitScmCleanBeforeCheckoutExtension),
				),
			},
			{
				Config: testAccJobGitScmConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmExtensions(&jobRef, []string{}, &extensions, ensureJobGitScmCleanBeforeCheckoutExtension),
				),
			},
		},
	})
}

func testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName string) string {
	return testAccJobGitScmConfigBasic(jobName) + `
resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
	scm = "${jenkins_job_git_scm.main.id}"
}`
}

func ensureJobGitScmCleanBeforeCheckoutExtension(
	extensionInterface client.GitScmExtension,
	rs *terraform.ResourceState,
) error {
	extension := extensionInterface.(*client.GitScmCleanBeforeCheckoutExtension)

	_, _, extensionId, err := resourceJobGitScmExtensionId(rs.Primary.Attributes["id"])
	if err != nil {
		return err
	}
	if extension.Id != extensionId {
		return fmt.Errorf("GitScmCleanBeforeCheckoutExtension id should be %v, was %v", extensionId, extension.Id)
	}

	return nil
}
