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

func TestAccJobGerritBranchBasic(t *testing.T) {
	var jobRef client.Job
	var branches []*client.JobGerritTriggerBranch
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	branchResourceName := "jenkins_job_gerrit_branch.main"
	compareType := "PLAIN"
	newCompareType := "REG_EXP"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritBranchConfigBasic(jobName, compareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(branchResourceName, "compare_type", compareType),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritBranches(&jobRef, []string{
						branchResourceName,
					}, &branches),
				),
			},
			{
				Config: testAccJobGerritBranchConfigBasic(jobName, newCompareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(branchResourceName, "compare_type", newCompareType),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritBranches(&jobRef, []string{
						branchResourceName,
					}, &branches),
				),
			},
			{
				ResourceName:  branchResourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid gerrit branch id"),
			},
			{
				ResourceName: branchResourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(branches) == 0 {
						return "", fmt.Errorf("no branches to import")
					}
					property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
					trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)
					project := (*trigger.Projects.Items)[0]
					return strings.Join([]string{jobName, property.GetId(), trigger.GetId(), project.Id, branches[0].Id}, IdDelimiter), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobGerritProjectConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritBranches(&jobRef, []string{}, &branches),
				),
			},
		},
	})
}

func testAccJobGerritBranchConfigBasic(jobName string, compareType string) string {
	return testAccJobGerritProjectConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_gerrit_branch" "main" {
  project = "${jenkins_job_gerrit_project.main.id}"

  compare_type = "%v"
  pattern      = "^(?!refs/meta/config).*$"
}
`, compareType)
}

func testAccCheckJobGerritBranches(jobRef *client.Job, expectedBranchResourceNames []string, returnBranches *[]*client.JobGerritTriggerBranch) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
		trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)
		project := (*trigger.Projects.Items)[0]

		if project.Branches.Items == nil && len(expectedBranchResourceNames) > 0 {
			return fmt.Errorf("Expected %v branches, found 0", len(expectedBranchResourceNames))
		}
		if project.Branches.Items != nil && len(*project.Branches.Items) != len(expectedBranchResourceNames) {
			return fmt.Errorf("Expected %v branches, found %v", len(expectedBranchResourceNames), len(*jobRef.Properties.Items))
		}
		for _, resourceName := range expectedBranchResourceNames {
			branchResource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Gerrit Branch Resource not found: %s", resourceName)
			}

			branch, err := ensureBranch(project, branchResource)
			if err != nil {
				return err
			}
			*returnBranches = append(*returnBranches, branch)
		}

		return nil
	}
}

func ensureBranch(project *client.JobGerritTriggerProject, resource *terraform.ResourceState) (*client.JobGerritTriggerBranch, error) {
	_, _, _, _, branchId, err := resourceJobGerritBranchId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	branch, err := project.GetBranch(branchId)
	if err != nil {
		return nil, err
	}

	if branch.CompareType.String() != resource.Primary.Attributes["compare_type"] {
		return nil, fmt.Errorf("expected compare_type %s, got %s",
			resource.Primary.Attributes["compare_type"], branch.CompareType)
	}
	if branch.Pattern != resource.Primary.Attributes["pattern"] {
		return nil, fmt.Errorf("expected pattern %s, got %s",
			resource.Primary.Attributes["pattern"], branch.Pattern)
	}

	return branch, nil
}
