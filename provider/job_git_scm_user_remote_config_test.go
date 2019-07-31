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
	var configs []*client.GitUserRemoteConfig
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	configResourceName := "jenkins_job_git_scm_user_remote_config.main"
	refspec := "my-refspec"
	newRefspec := "new-my-refspec"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, refspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(configResourceName, "refspec", refspec),
					resource.TestCheckResourceAttr(configResourceName, "url", "my-test-url"),
					resource.TestCheckResourceAttr(configResourceName, "credentials_id", "my-test-creds"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmUserRemoteConfigs(&jobRef, []string{
						configResourceName,
					}, &configs),
				),
			},
			{
				Config: testAccJobGitScmUserRemoteConfigConfigBasic(jobName, newRefspec),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(configResourceName, "refspec", newRefspec),
					resource.TestCheckResourceAttr(configResourceName, "url", "my-test-url"),
					resource.TestCheckResourceAttr(configResourceName, "credentials_id", "my-test-creds"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmUserRemoteConfigs(&jobRef, []string{
						configResourceName,
					}, &configs),
				),
			},
			{
				Config: testAccJobGitScmConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmUserRemoteConfigs(&jobRef, []string{}, &configs),
				),
			},
		},
	})
}

func testAccJobGitScmUserRemoteConfigConfigBasic(jobName string, refspec string) string {
	return testAccJobGitScmConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_git_scm_user_remote_config" "main" {
	scm = "${jenkins_job_git_scm.main.id}"

  refspec        = "%s"
  url            = "my-test-url"
  credentials_id = "my-test-creds"
}`, refspec)
}

func testAccCheckJobGitScmUserRemoteConfigs(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnConfigs *[]*client.GitUserRemoteConfig,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		definition := jobRef.Definition.(*client.CpsScmFlowDefinition)

		if len(expectedResourceNames) == 0 && definition.SCM.UserRemoteConfigs.Items == nil {
			return nil
		}
		if definition.SCM.UserRemoteConfigs.Items == nil {
			return fmt.Errorf("Expected %v git scm user remote configs, found %v",
				len(expectedResourceNames), 0)
		}
		if len(*definition.SCM.UserRemoteConfigs.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v git scm user remote configs, found %v",
				len(expectedResourceNames), len(*definition.SCM.UserRemoteConfigs.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Git Scm User Remote Config Resource not found: %s", resourceName)
			}

			_, _, configId, err := resourceJobGitScmUserRemoteConfigId(resource.Primary.Attributes["id"])
			config, err := definition.SCM.GetUserRemoteConfig(configId)
			if err != nil {
				return err
			}
			err = ensureJobGitScmUserRemoteConfig(config, resource)
			if err != nil {
				return err
			}
			*returnConfigs = append(*returnConfigs, config)
		}

		return nil
	}
}

func ensureJobGitScmUserRemoteConfig(
	config *client.GitUserRemoteConfig,
	resource *terraform.ResourceState,
) error {
	if config.Refspec != resource.Primary.Attributes["refspec"] {
		return fmt.Errorf("expected refspec %s, got %s",
			resource.Primary.Attributes["refspec"], config.Refspec)
	}
	if config.Url != resource.Primary.Attributes["url"] {
		return fmt.Errorf("expected url %s, got %s",
			resource.Primary.Attributes["url"], config.Url)
	}
	if config.CredentialsId != resource.Primary.Attributes["credentials_id"] {
		return fmt.Errorf("expected credentials_id %s, got %s",
			resource.Primary.Attributes["credentials_id"], config.CredentialsId)
	}

	return nil
}
