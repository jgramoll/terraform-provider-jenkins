package provider

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func testAccCheckJobDefinition(jobRef *client.Job, expectedDefinitionResourceName string, returnDefinition *client.JobDefinition, ensureJobDefinitionFunc func(client.JobDefinition, *terraform.ResourceState) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		definitionResource, ok := s.RootModule().Resources[expectedDefinitionResourceName]
		if !ok {
			return fmt.Errorf("Job Defintion Resource not found: %s", expectedDefinitionResourceName)
		}

		*returnDefinition = jobRef.Definition
		return ensureJobDefinitionFunc(jobRef.Definition, definitionResource)
	}
}

func testAccCheckNoJobDefinition(jobRef *client.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if jobRef.Definition != nil {
			return errors.New("Job should not have definition")
		}
		return nil
	}
}
