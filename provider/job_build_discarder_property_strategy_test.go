package provider

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobBuildDiscarderPropertyStrategyTypes = map[string]reflect.Type{}

func testAccCheckBuildDiscarderPropertyStrategies(jobRef *client.Job, expectedStrategyResourceNames []string, returnedStrategies *[]client.JobBuildDiscarderPropertyStrategy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if jobRef.Properties == nil {
			return fmt.Errorf("Unexpected nil properties")
		}

		if len(*jobRef.Properties.Items) != len(expectedStrategyResourceNames) {
			return fmt.Errorf("Expected %v properties, found %v", len(expectedStrategyResourceNames), len(*jobRef.Properties.Items))
		}
		for _, resourceName := range expectedStrategyResourceNames {
			strategyResource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Property Resource not found: %s", resourceName)
			}

			strategy, err := ensureJobBuildDiscarderPropertyStrategy(jobRef, strategyResource)
			if err != nil {
				return err
			}
			*returnedStrategies = append(*returnedStrategies, strategy)
		}

		return nil
	}
}

func ensureJobBuildDiscarderPropertyStrategy(jobRef *client.Job, resource *terraform.ResourceState) (client.JobBuildDiscarderPropertyStrategy, error) {
	_, propertyId, strategyId, err := resourceJobPropertyStrategyId(resource.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	property, err := jobRef.GetProperty(propertyId)
	if err != nil {
		return nil, err
	}

	buildDiscarderProperty := property.(*client.JobBuildDiscarderProperty)
	strategy := buildDiscarderProperty.Strategy.Item

	if strategy == nil {
		return nil, errors.New("Property does not have strategy")
	}

	expectedType := jobBuildDiscarderPropertyStrategyTypes[resource.Type]
	strategyType := reflect.TypeOf(strategy)
	if expectedType != strategyType {
		return nil, fmt.Errorf("Job Property Strategy %s was type \"%s\", expected type \"%s\"",
			strategyId, strategyType, expectedType)
	}
	return strategy, nil
}
