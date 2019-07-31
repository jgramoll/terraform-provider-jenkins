package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobTriggerTypes = map[string]reflect.Type{}

func testAccCheckJobTriggers(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnTriggers *[]client.JobTrigger,
	ensureTrigger func(client.JobTrigger, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
		if len(*property.Triggers.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v triggers, found %v", len(expectedResourceNames), len(*property.Triggers.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Trigger Resource not found: %s", resourceName)
			}

			_, _, triggerId, err := resourceJobTriggerId(resource.Primary.Attributes["id"])
			if err != nil {
				return err
			}

			triggerInteface, err := property.GetTrigger(triggerId)
			if err != nil {
				return err
			}

			err = ensureTrigger(triggerInteface, resource)
			if err != nil {
				return err
			}
			*returnTriggers = append(*returnTriggers, triggerInteface)
		}

		return nil
	}
}
