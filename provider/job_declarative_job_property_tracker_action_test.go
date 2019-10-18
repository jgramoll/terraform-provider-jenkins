package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobDeclarativeJobPropertyTrackerActionBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDeclarativeJobPropertyTrackerActionConfig(jobName, "pipeline-model-definition@1.3.7"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "action.0.type", "DeclarativeJobPropertyTrackerAction"),
					resource.TestCheckResourceAttr(jobResourceName, "action.0.plugin", "pipeline-model-definition@1.3.7"),
				),
			},
			{
				Config: testAccJobDeclarativeJobPropertyTrackerActionConfig(jobName, "pipeline-model-definition@1.3.8"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "action.0.type", "DeclarativeJobPropertyTrackerAction"),
					resource.TestCheckResourceAttr(jobResourceName, "action.0.plugin", "pipeline-model-definition@1.3.8"),
				),
			},
		},
	})
}

func testAccJobDeclarativeJobPropertyTrackerActionConfig(jobName string, plugin string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name   = "%s"
	action {
		type = "DeclarativeJobPropertyTrackerAction"
		plugin = "%s"
	}
}`, jobName, plugin)
}
