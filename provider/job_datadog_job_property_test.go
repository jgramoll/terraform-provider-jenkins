package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobDatadogJobPropertyBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDatadogJobPropertyConfigBasic(jobName, "datadog@0.7.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.type", "DatadogJobProperty"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.plugin", "datadog@0.7.0"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.emit_on_checkout", "false"),
				),
			},
			{
				Config: testAccJobDatadogJobPropertyConfigBasic(jobName, "datadog@0.7.1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.type", "DatadogJobProperty"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.plugin", "datadog@0.7.1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.emit_on_checkout", "false"),
				),
			},
		},
	})
}

func testAccJobDatadogJobPropertyConfigBasic(jobName string, plugin string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

	property {
		type   = "DatadogJobProperty"
		plugin = "%s"

		emit_on_checkout = false
	}
}`, jobName, plugin)
}
