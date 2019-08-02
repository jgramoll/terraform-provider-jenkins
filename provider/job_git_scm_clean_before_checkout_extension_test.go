package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

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
				ResourceName:  extensionResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid git scm extension id"),
			},
			{
				ResourceName: extensionResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(extensions) == 0 {
						return "", fmt.Errorf("no extensions to import")
					}
					definitionId := jobRef.Definition.GetId()
					return strings.Join([]string{jobName, definitionId, extensions[0].GetId()}, IdDelimiter), nil
				},
				ImportStateVerify: true,
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
	_, err := newJobGitScmCleanBeforeCheckoutExtension().fromClientExtension(extensionInterface)
	return err
}
