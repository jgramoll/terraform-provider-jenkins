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

func TestAccJobGerritFilePathBasic(t *testing.T) {
	var jobRef client.Job
	var resources []*client.JobGerritTriggerFilePath
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	resourceName := "jenkins_job_gerrit_file_path.main"
	compareType := "PLAIN"
	newCompareType := "REG_EXP"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritFilePathConfigBasic(jobName, compareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compare_type", compareType),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritFilePaths(&jobRef, []string{
						resourceName,
					}, &resources),
				),
			},
			{
				Config: testAccJobGerritFilePathConfigBasic(jobName, newCompareType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compare_type", newCompareType),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritFilePaths(&jobRef, []string{
						resourceName,
					}, &resources),
				),
			},
			{
				ResourceName:  resourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid gerrit file path id"),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(resources) == 0 {
						return "", fmt.Errorf("no file paths to import")
					}
					property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
					trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)
					project := (*trigger.Projects.Items)[0]
					return ResourceJobGerritFilePathId(jobName, property.GetId(), trigger.GetId(), project.Id, resources[0].Id), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobGerritProjectConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritFilePaths(&jobRef, []string{}, &resources),
				),
			},
		},
	})
}

func testAccJobGerritFilePathConfigBasic(jobName string, compareType string) string {
	return testAccJobGerritProjectConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_gerrit_file_path" "main" {
  project = "${jenkins_job_gerrit_project.main.id}"

  compare_type = "%v"
  pattern      = "^(?!refs/meta/config).*$"
}
`, compareType)
}

func testAccCheckJobGerritFilePaths(
	jobRef *client.Job, expectedResourceNames []string,
	returnResources *[]*client.JobGerritTriggerFilePath,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
		trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)
		project := (*trigger.Projects.Items)[0]

		if project.FilePaths.Items == nil && len(expectedResourceNames) > 0 {
			return fmt.Errorf("Expected %v files paths, found 0", len(expectedResourceNames))
		}
		if project.FilePaths.Items != nil && len(*project.FilePaths.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v file paths, found %v", len(expectedResourceNames), len(*jobRef.Properties.Items))
		}
		for _, resourceName := range expectedResourceNames {
			filePathResource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Gerrit FilePath Resource not found: %s", resourceName)
			}

			filePath, err := ensureFilePath(project, filePathResource)
			if err != nil {
				return err
			}
			*returnResources = append(*returnResources, filePath)
		}

		return nil
	}
}

func ensureFilePath(
	project *client.JobGerritTriggerProject, resource *terraform.ResourceState,
) (*client.JobGerritTriggerFilePath, error) {
	_, _, _, _, filePathId, err := resourceJobGerritFilePathParseId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	filePath, err := project.GetFilePath(filePathId)
	if err != nil {
		return nil, err
	}

	if filePath.CompareType.String() != resource.Primary.Attributes["compare_type"] {
		return nil, fmt.Errorf("expected compare_type %s, got %s",
			resource.Primary.Attributes["compare_type"], filePath.CompareType)
	}
	if filePath.Pattern != resource.Primary.Attributes["pattern"] {
		return nil, fmt.Errorf("expected pattern %s, got %s",
			resource.Primary.Attributes["pattern"], filePath.Pattern)
	}

	return filePath, nil
}
