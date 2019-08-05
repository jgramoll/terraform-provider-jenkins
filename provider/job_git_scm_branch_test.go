package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGitBranchBasic(t *testing.T) {
	var jobRef client.Job
	var branches []*client.GitScmBranchSpec
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	branchResourceName := "jenkins_job_git_scm_branch.main"
	branchName := "my-branch"
	newBranchName := "my-new-branch"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGitScmBranchConfigBasic(jobName, branchName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(branchResourceName, "name", branchName),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmBranches(&jobRef, []string{
						branchResourceName,
					}, &branches),
				),
			},
			{
				Config: testAccJobGitScmBranchConfigBasic(jobName, newBranchName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(branchResourceName, "name", newBranchName),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGitScmBranches(&jobRef, []string{
						branchResourceName,
					}, &branches),
				),
			},
			{
				ResourceName:  branchResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid git scm branch id"),
			},
			{
				ResourceName: branchResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(branches) == 0 {
						return "", fmt.Errorf("no branches to import")
					}
					definitionId := jobRef.Definition.GetId()
					return ResourceJobGitScmBranchId(jobName, definitionId, branches[0].Id), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobGitScmConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
				),
			},
		},
	})
}

func testAccJobGitScmBranchConfigBasic(jobName string, branchName string) string {
	return testAccJobGitScmConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_git_scm_branch" "main" {
  scm = "${jenkins_job_git_scm.main.id}"

  name = "%s"
}`, branchName)
}

func testAccCheckJobGitScmBranches(jobRef *client.Job, expectedResourceNames []string, returnBranches *[]*client.GitScmBranchSpec) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		definition := jobRef.Definition.(*client.CpsScmFlowDefinition)

		if len(*definition.SCM.Branches.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v branches, found %v", len(expectedResourceNames), len(*definition.SCM.Branches.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Git Scm Branch Resource not found: %s", resourceName)
			}

			branch, err := ensureGitScmBranch(definition, resource)
			if err != nil {
				return err
			}
			*returnBranches = append(*returnBranches, branch)
		}

		return nil
	}
}

func ensureGitScmBranch(definition *client.CpsScmFlowDefinition, resource *terraform.ResourceState) (*client.GitScmBranchSpec, error) {
	_, _, branchId, err := resourceJobGitScmBranchParseId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	branch, err := definition.SCM.GetBranch(branchId)
	if err != nil {
		return nil, err
	}

	if branch.Name != resource.Primary.Attributes["name"] {
		return nil, fmt.Errorf("expected name %s, got %s",
			resource.Primary.Attributes["name"], branch.Name)
	}

	return branch, nil
}
