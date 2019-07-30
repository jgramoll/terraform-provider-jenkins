package provider

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobDefinitionTypes = map[string]reflect.Type{}

func testAccCheckJobDefinition(jobRef *client.Job, expectedDefinitionResourceName string, ensureJobDefinitionFunc func(client.JobDefinition, *terraform.ResourceState) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		definitionResource, ok := s.RootModule().Resources[expectedDefinitionResourceName]
		if !ok {
			return fmt.Errorf("Job Defintion Resource not found: %s", expectedDefinitionResourceName)
		}

		_, err := ensureJobDefinition(jobRef, definitionResource, ensureJobDefinitionFunc)
		if err != nil {
			return err
		}
		return nil
	}
}

func ensureJobDefinition(jobRef *client.Job, rs *terraform.ResourceState, ensureJobDefinitionFunc func(client.JobDefinition, *terraform.ResourceState) error) (client.JobDefinition, error) {
	_, definitionId, err := resourceJobDefinitionId(rs.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	expectedType := jobDefinitionTypes[rs.Type]
	definitionType := reflect.TypeOf(jobRef.Definition)
	if expectedType != definitionType {
		return nil, fmt.Errorf("Job Defintion %s was type \"%s\", expected type \"%s\"",
			definitionId, definitionType, expectedType)
	}

	err = ensureJobDefinitionFunc(jobRef.Definition, rs)
	if err != nil {
		return nil, err
	}

	return jobRef.Definition, nil
}

func testAccCheckNoJobDefinition(jobRef *client.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if jobRef.Definition != nil {
			return errors.New("Job should not have definition")
		}
		return nil
	}
}
