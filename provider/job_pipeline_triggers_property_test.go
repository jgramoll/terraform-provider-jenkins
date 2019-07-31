package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyTypes["jenkins_job_pipeline_triggers_property"] = reflect.TypeOf((*client.JobPipelineTriggersProperty)(nil))
}

func TestAccJobPipelineTriggersPropertyBasic(t *testing.T) {
	var jobRef client.Job
	var properties []client.JobProperty
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	property1 := "jenkins_job_pipeline_triggers_property.prop_1"
	property2 := "jenkins_job_pipeline_triggers_property.prop_2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobPipelineTriggersPropertyConfigBasic(jobName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
						property2,
					}, &properties, ensureJobTriggersProperty),
				),
			},
			{
				Config: testAccJobPipelineTriggersPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
					}, &properties, ensureJobTriggersProperty),
				),
			},
			{
				Config: testAccJobPipelineTriggersPropertyConfigBasic(jobName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{}, &properties, ensureJobTriggersProperty),
				),
			},
		},
	})
}

func testAccJobPipelineTriggersPropertyConfigBasic(jobName string, count int) string {
	properties := ""
	for i := 1; i <= count; i++ {
		properties += fmt.Sprintf(`
resource "jenkins_job_pipeline_triggers_property" "prop_%v" {
	job = "${jenkins_job.main.id}"
}`, i)
	}

	return testAccJobConfigBasic(jobName) + properties
}

func ensureJobTriggersProperty(propertyInterface client.JobProperty, resource *terraform.ResourceState) error {
	_, ok := propertyInterface.(*client.JobPipelineTriggersProperty)
	if !ok {
		return fmt.Errorf("Property is not of expected type, expected *client.JobPipelineTriggersProperty, actually %s",
			reflect.TypeOf(propertyInterface).String())
	}
	return nil
}
