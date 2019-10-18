package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobPipelineTriggersPropertyBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobPipelineTriggersPropertyConfigBasic(jobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.type", "PipelineTriggersJobProperty"),
				),
			},
		},
	})
}

func testAccJobPipelineTriggersPropertyConfigBasic(jobName string) string {
	return fmt.Sprintf(`

	resource "jenkins_job" "main" {
		name     = "%s"
		plugin   = "workflow-job@2.33"
	
		property {
			type = "PipelineTriggersJobProperty"
		}
	}`, jobName)
}
