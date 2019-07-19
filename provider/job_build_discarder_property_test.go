package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyTypes["jenkins_job_build_discarder_property"] = reflect.TypeOf((*client.JobBuildDiscarderProperty)(nil))
}

func TestAccJobBuildDiscarderPropertyBasic(t *testing.T) {
	var jobRef client.Job
	var properties []client.JobProperty

	jobName := fmt.Sprintf("Bridge Career/tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	property1 := "jenkins_job_build_discarder_property.prop_1"
	property2 := "jenkins_job_build_discarder_property.prop_2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobBuildDiscarderPropertyConfigBasic(jobName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
						property2,
					}, &properties),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
					}, &properties),
				),
			},
			{
				Config: testAccJobBuildDiscarderPropertyConfigBasic(jobName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{}, &properties),
				),
			},
		},
	})
}

func testAccJobBuildDiscarderPropertyConfigBasic(jobName string, count int) string {
	properties := ""
	for i := 1; i <= count; i++ {
		properties += fmt.Sprintf(`
resource "jenkins_job_build_discarder_property" "prop_%v" {
	job = "${jenkins_job.main.id}"
}`, i)
	}

	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name = "%s"
}`, jobName) + properties
}
