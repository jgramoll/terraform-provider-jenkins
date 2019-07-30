package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobGerritProjectBasic(t *testing.T) {
	var jobRef client.Job
	var projects []*client.JobGerritTriggerProject
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	projectResourceName := "jenkins_job_gerrit_project.main"
	pattern := "bridge-skills"
	newPattern := "new-bridge-skills"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritProjectConfigPattern(jobName, pattern),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "pattern", pattern),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritProjects(&jobRef, []string{
						projectResourceName,
					}, &projects),
				),
			},
			{
				Config: testAccJobGerritProjectConfigPattern(jobName, newPattern),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "pattern", newPattern),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritProjects(&jobRef, []string{
						projectResourceName,
					}, &projects),
				),
			},
			{
				Config: testAccJobGerritTriggerConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobGerritProjects(&jobRef, []string{}, &projects),
				),
			},
		},
	})
}

func testAccJobGerritProjectConfigBasic(jobName string) string {
	return testAccJobGerritProjectConfigPattern(jobName, "bridge-skills")
}

func testAccJobGerritProjectConfigPattern(jobName string, pattern string) string {
	return testAccJobGerritTriggerConfigBasic(jobName) + fmt.Sprintf(`
resource "jenkins_job_gerrit_project" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1.id}"

	compare_type = "PLAIN"
	pattern      = "%v"
}
`, pattern)
}

func testAccCheckJobGerritProjects(jobRef *client.Job, expectedResourceNames []string, returnProjects *[]*client.JobGerritTriggerProject) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
		trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)

		if len(*trigger.Projects.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v projects, found %v", len(expectedResourceNames), len(*jobRef.Properties.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Gerrit Project Resource not found: %s", resourceName)
			}

			project, err := ensureProject(trigger, resource)
			if err != nil {
				return err
			}
			*returnProjects = append(*returnProjects, project)
		}

		return nil
	}
}

func ensureProject(trigger *client.JobGerritTrigger, resource *terraform.ResourceState) (*client.JobGerritTriggerProject, error) {
	_, _, _, projectId, err := resourceJobGerritProjectId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	project, err := trigger.GetProject(projectId)
	if err != nil {
		return nil, err
	}

	if project.CompareType.String() != resource.Primary.Attributes["compare_type"] {
		return nil, fmt.Errorf("expected compare_type %s, got %s",
			resource.Primary.Attributes["compare_type"], project.CompareType)
	}
	if project.Pattern != resource.Primary.Attributes["pattern"] {
		return nil, fmt.Errorf("expected pattern %s, got %s",
			resource.Primary.Attributes["pattern"], project.Pattern)
	}

	return project, nil
}
