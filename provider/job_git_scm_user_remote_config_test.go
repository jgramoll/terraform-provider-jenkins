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
	jobResourceName := "jenkins_job.main"
	refspec := "my-refspec"
	newRefspec := "new-my-refspec"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, refspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.refspec", refspec),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.url", "my-test-url"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.credentials_id", "my-test-creds"),
				),
			},
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, newRefspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.refspec", newRefspec),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.url", "my-test-url"),
					resource.TestCheckResourceAttr(jobResourceName, "definition.0.scm.0.user_remote_config.0.credentials_id", "my-test-creds"),
				),
			},
		},
	})
}

func testAccJobGitScmUserRemoteConfigConfigBasic(jobName string, refspec string) string {
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

			user_remote_config {
				refspec        = "%s"
				url            = "my-test-url"
				credentials_id = "my-test-creds"
			}
		}
	}
}`, jobName, refspec)
}
