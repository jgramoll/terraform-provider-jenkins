package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func testAccCheckJobParameterDefintion(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnItems *[]client.JobParameterDefinition,
	ensureFunc func(client.JobParameterDefinition, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Property Resource not found: %s", resourceName)
			}

			_, propertyId, definitionId, err := resourceJobParameterDefinitionParseId(resource.Primary.Attributes["id"])
			if err != nil {
				return err
			}

			propertyInterface, err := jobRef.GetProperty(propertyId)
			if err != nil {
				return err
			}
			property, ok := propertyInterface.(*client.JobParametersDefinitionProperty)
			if !ok {
				return fmt.Errorf("Could not get parameters definition property from %v", property)
			}
			definition, err := property.GetParameterDefinition(definitionId)
			if err != nil {
				return err
			}

			err = ensureFunc(definition, resource)
			if err != nil {
				return err
			}

			*returnItems = append(*returnItems, definition)
		}

		return nil
	}
}
