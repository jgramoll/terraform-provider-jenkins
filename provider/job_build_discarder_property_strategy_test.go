package provider

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func testAccCheckBuildDiscarderPropertyStrategies(
	jobRef *client.Job,
	expectedStrategyResourceNames []string,
	returnedStrategies *[]client.JobBuildDiscarderPropertyStrategy,
	ensureStrategyFunc func(client.JobBuildDiscarderPropertyStrategy, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if jobRef.Properties == nil {
			return fmt.Errorf("Unexpected nil properties")
		}

		strategyCount := 0
		for _, p := range *(jobRef.Properties.Items) {
			buildDiscarderProperty := p.(*client.JobBuildDiscarderProperty)
			if buildDiscarderProperty.Strategy.Item != nil {
				strategyCount += 1
			}
		}
		if strategyCount != len(expectedStrategyResourceNames) {
			return fmt.Errorf("Expected %v Job Strategy Resources, found %v", expectedStrategyResourceNames, strategyCount)
		}

		for _, resourceName := range expectedStrategyResourceNames {
			strategyResource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Property Resource not found: %s", resourceName)
			}

			strategy, err := ensureJobBuildDiscarderPropertyStrategy(jobRef, strategyResource, ensureStrategyFunc)
			if err != nil {
				return err
			}
			*returnedStrategies = append(*returnedStrategies, strategy)
		}

		return nil
	}
}

func ensureJobBuildDiscarderPropertyStrategy(
	jobRef *client.Job,
	resource *terraform.ResourceState,
	ensureStrategyFunc func(client.JobBuildDiscarderPropertyStrategy, *terraform.ResourceState) error,
) (client.JobBuildDiscarderPropertyStrategy, error) {
	_, propertyId, _, err := resourceJobPropertyStrategyId(resource.Primary.Attributes["id"])
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

	err = ensureStrategyFunc(strategy, resource)
	if err != nil {
		return nil, err
	}

	return strategy, nil
}
