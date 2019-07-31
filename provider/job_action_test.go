package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func testAccCheckJobActions(
	jobRef *client.Job,
	expectedResourceNames []string,
	returnActions *[]client.JobAction,
	ensureAction func(client.JobAction, *terraform.ResourceState) error,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if len(*jobRef.Actions.Items) != len(expectedResourceNames) {
			return fmt.Errorf("Expected %v actions, found %v", len(expectedResourceNames), len(*jobRef.Actions.Items))
		}
		for _, resourceName := range expectedResourceNames {
			resource, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Job Action Resource not found: %s", resourceName)
			}

			_, actionId, err := resourceJobActionId(resource.Primary.Attributes["id"])
			if err != nil {
				return err
			}

			actionInterface, err := jobRef.GetAction(actionId)
			if err != nil {
				return err
			}

			err = ensureAction(actionInterface, resource)
			if err != nil {
				return err
			}
			*returnActions = append(*returnActions, actionInterface)
		}

		return nil
	}
}
