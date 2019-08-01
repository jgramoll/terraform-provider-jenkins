package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobDatadogJobPropertyBasic(t *testing.T) {
	var jobRef client.Job
	var properties []client.JobProperty
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	property1 := "jenkins_job_datadog_job_property.prop_1"
	property2 := "jenkins_job_datadog_job_property.prop_2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDatadogJobPropertyConfigBasic(jobName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
						property2,
					}, &properties, ensureJobDatadogJobProperty),
				),
			},
			{
				Config: testAccJobDatadogJobPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
					}, &properties, ensureJobDatadogJobProperty),
				),
			},
			{
				Config: testAccJobDatadogJobPropertyConfigBasic(jobName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{}, &properties, ensureJobDatadogJobProperty),
				),
			},
		},
	})
}

func testAccJobDatadogJobPropertyConfigBasic(jobName string, count int) string {
	properties := ""
	for i := 1; i <= count; i++ {
		properties += fmt.Sprintf(`
resource "jenkins_job_datadog_job_property" "prop_%v" {
	job = "${jenkins_job.main.id}"
}`, i)
	}

	return testAccJobConfigBasic(jobName) + properties
}

func ensureJobDatadogJobProperty(propertyInterface client.JobProperty, resource *terraform.ResourceState) error {
	_, err := newJobDatadogJobProperty().fromClientJobProperty(propertyInterface)
	return err
}
