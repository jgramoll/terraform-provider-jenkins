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

func TestAccJobDisableConcurrentBuildsJobPropertyBasic(t *testing.T) {
	var jobRef client.Job
	var properties []client.JobProperty
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	property1 := "jenkins_job_disable_concurrent_builds_property.prop_1"
	property2 := "jenkins_job_disable_concurrent_builds_property.prop_2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobDisableConcurrentBuildsJobPropertyConfigBasic(jobName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
						property2,
					}, &properties, ensureJobDisableConcurrentBuildsJobProperty),
				),
			},
			{
				ResourceName:  property1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid property id"),
			},
			{
				ResourceName: property1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(properties) == 0 {
						return "", fmt.Errorf("no properties to import")
					}
					return ResourceJobPropertyId(jobName, properties[0].GetId()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: property2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(properties) == 0 {
						return "", fmt.Errorf("no properties to import")
					}
					return ResourceJobPropertyId(jobName, properties[1].GetId()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobDisableConcurrentBuildsJobPropertyConfigBasic(jobName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{
						property1,
					}, &properties, ensureJobDisableConcurrentBuildsJobProperty),
				),
			},
			{
				Config: testAccJobDisableConcurrentBuildsJobPropertyConfigBasic(jobName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobProperties(&jobRef, []string{}, &properties, ensureJobDisableConcurrentBuildsJobProperty),
				),
			},
		},
	})
}

func testAccJobDisableConcurrentBuildsJobPropertyConfigBasic(jobName string, count int) string {
	properties := ""
	for i := 1; i <= count; i++ {
		properties += fmt.Sprintf(`
resource "jenkins_job_disable_concurrent_builds_property" "prop_%v" {
	job = "${jenkins_job.main.name}"
}`, i)
	}

	return testAccJobConfigBasic(jobName) + properties
}

func ensureJobDisableConcurrentBuildsJobProperty(propertyInterface client.JobProperty, resource *terraform.ResourceState) error {
	_, err := newJobDisableConcurrentBuildsJobProperty().fromClientJobProperty(propertyInterface)
	return err
}
