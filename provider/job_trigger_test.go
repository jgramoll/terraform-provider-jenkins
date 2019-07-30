package provider

import (
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobTriggerTypes = map[string]reflect.Type{}

func testAccCheckJobTriggers(jobRef *client.Job, expectedResourceNames []string, returnTriggers *[]client.JobTrigger) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// if len(*jobRef.Properties.Items) != len(expectedPropertyResourceNames) {
		// 	return fmt.Errorf("Expected %v properties, found %v", len(expectedPropertyResourceNames), len(*jobRef.Properties.Items))
		// }
		// for _, resourceName := range expectedPropertyResourceNames {
		// 	propertyResource, ok := s.RootModule().Resources[resourceName]
		// 	if !ok {
		// 		return fmt.Errorf("Job Property Resource not found: %s", resourceName)
		// 	}

		// 	property, err := ensureProperty(jobRef, propertyResource)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	*returnProperties = append(*returnProperties, property)
		// }

		return nil
	}
}

func ensureTrigger(jobRef *client.Job, resource *terraform.ResourceState) (client.JobTrigger, error) {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}
	println("jobName", jobName)
	println("triggerId", triggerId)

	property, err := jobRef.GetProperty(propertyId)
	if err != nil {
		return nil, err
	}

	// jobAttribute := resource.Primary.Attributes["job"]
	// if jobName != jobAttribute {
	// 	return nil, fmt.Errorf("Property Job should be %s, was %s", jobName, jobAttribute)
	// }

	// expectedType := jobPropertyTypes[resource.Type]
	// propertyType := reflect.TypeOf(property)
	// if expectedType != propertyType {
	// 	return nil, fmt.Errorf("Job Property %s was type \"%s\", expected type \"%s\"",
	// 		propertyId, propertyType, expectedType)
	// }
	return property, nil
}
