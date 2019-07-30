package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobTriggerEventTypes = map[string]reflect.Type{}

func testAccCheckJobGerritTriggerEvents(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnEvents *[]client.JobGerritTriggerOnEvent,
	ensureTriggerEventFunc func(client.JobGerritTriggerOnEvent, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		property := (*jobRef.Properties.Items)[0].(*client.JobPipelineTriggersProperty)
		trigger := (*property.Triggers.Items)[0].(*client.JobGerritTrigger)

		if len(*trigger.TriggerOnEvents.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v trigger events, found %v",
				len(expectedResourceNames), len(*trigger.TriggerOnEvents.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Trigger Event Resource not found: %s", resourceName)
			}

			event, err := ensureTriggerEvent(jobRef, resource, ensureTriggerEventFunc)
			if err != nil {
				return err
			}
			*returnEvents = append(*returnEvents, event)
		}

		return nil
	}
}

func ensureTriggerEvent(
	jobRef *client.Job,
	resource *terraform.ResourceState,
	ensureTriggerEventFunc func(client.JobGerritTriggerOnEvent, *terraform.ResourceState) error,
) (client.JobTrigger, error) {
	_, propertyId, triggerId, eventId, err := resourceJobTriggerEventId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	propertyInterface, err := jobRef.GetProperty(propertyId)
	if err != nil {
		return nil, err
	}
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	triggerInterface, err := property.GetTrigger(triggerId)
	if err != nil {
		return nil, err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	event, err := trigger.GetEvent(eventId)
	if err != nil {
		return nil, err
	}

	expectedType := jobTriggerEventTypes[resource.Type]
	eventType := reflect.TypeOf(event)
	if expectedType != eventType {
		return nil, fmt.Errorf("Job Event %s was type \"%s\", expected type \"%s\"",
			propertyId, eventType, expectedType)
	}

	err = ensureTriggerEventFunc(event, resource)
	if err != nil {
		return nil, err
	}

	return property, nil
}
