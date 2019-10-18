package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobJiraProjectPropertyBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobJiraProjectPropertyConfigBasic(jobName, "jira@3.0.7"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.type", "JiraProjectProperty"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.plugin", "jira@3.0.7"),
				),
			},
			{
				Config: testAccJobJiraProjectPropertyConfigBasic(jobName, "jira@3.0.8"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.type", "JiraProjectProperty"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.plugin", "jira@3.0.8"),
				),
			},
		},
	})
}

func testAccJobJiraProjectPropertyConfigBasic(jobName string, plugin string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

  property {
    type = "JiraProjectProperty"
    plugin="%s"
  }
}`, jobName, plugin)
}
