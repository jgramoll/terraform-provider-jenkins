package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobGitScmExtensionTypes = map[string]reflect.Type{}

func testAccCheckJobGitScmExtensions(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnEvents *[]client.GitScmExtension,
	ensureExtensionFunc func(client.GitScmExtension, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		definition := jobRef.Definition.(*client.CpsScmFlowDefinition)

		if len(*definition.SCM.Extensions.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v git scm extensions, found %v",
				len(expectedResourceNames), len(*definition.SCM.Extensions.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Trigger Event Resource not found: %s", resourceName)
			}

			_, _, extensionId, err := resourceJobGitScmExtensionId(resource.Primary.Attributes["id"])
			extension, err := definition.SCM.GetExtension(extensionId)
			if err != nil {
				return err
			}
			err = ensureExtensionFunc(extension, resource)
			if err != nil {
				return err
			}
			*returnEvents = append(*returnEvents, extension)
		}

		return nil
	}
}
