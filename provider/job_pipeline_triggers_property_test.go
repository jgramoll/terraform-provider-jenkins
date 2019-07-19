package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	// jobPropertyTypes["jenkins_job_pipeline_triggers_property"] = ""
}

func TestAccJobGerritPropertyBasic(t *testing.T) {
	var jobRef client.Job
	var properties []client.JobProperty
	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.test"
	property1 := "jenkins_job_pipeline_triggers_property.prop_1"
	property2 := "jenkins_job_pipeline_triggers_property.prop_2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobGerritPropertyConfigBasic(jobName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
						property2,
					}, &properties),
				),
			},
			{
				Config: testAccJobGerritPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
					}, &properties),
				),
			},
			{
				Config: testAccJobGerritPropertyConfigBasic(jobName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{}, &properties),
				),
			},
		},
	})
}

func testAccJobGerritPropertyConfigBasic(jobName string, count int) string {
	properties := ""
	for i := 1; i <= count; i++ {
		properties += fmt.Sprintf(`
resource "jenkins_job_pipeline_triggers_property" "prop_%v" {
	job = "${jenkins_job.test.id}"
}`, i)
	}

	return fmt.Sprintf(`
resource "jenkins_job" "test" {
	name = "%s"
}`, jobName) + properties
}
