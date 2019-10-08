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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.extension.0.type", "CleanBeforeCheckout"),
				),
			},
		},
	})
}

func testAccJobGitScmCleanBeforeCheckoutExtensionConfigBasic(jobName string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name   = "%s"
	plugin = "workflow-job@2.33"

	definition {
		type   = "CpsScmFlowDefinition"
		plugin = "workflow-cps@2.70"

		scm {
			type   = "GitSCM"
			plugin = "git@3.10.0"

			config_version = "2"

			extension {
				type = "CleanBeforeCheckout"
			}
		}
	}
}`, jobName)
}
