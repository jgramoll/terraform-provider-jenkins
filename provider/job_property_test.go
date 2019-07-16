package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

var jobPropertyTypes map[string]string // TODO TYPE?

func testAccCheckJobProperties(resourceName string, expected []string, properties *[]client.JobProperty) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Pipeline not found: %s", resourceName)
		}

		jobService := testAccProvider.Meta().(*Services).JobService
		job, err := jobService.GetJob(rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		properties := *(*job.Properties).Items
		if len(expected) != len(properties) {
			return fmt.Errorf("Job Property count of %v is expected to be %v",
				len(properties), len(expected))
		}

		for _, stageResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[stageResourceName]
			if !ok {
				return fmt.Errorf("Property not found: %s", resourceName)
			}
			println("expectedResource", expectedResource)

			// stage, err := ensureStage(pipeline, expectedResource)
			// if err != nil {
			// 	return err
			// }
			// *stages = append(*stages, stage)
		}

		return nil
	}
}

func testAccCheckJobPropertyDestroy(s *terraform.State) error {
	jobService := testAccProvider.Meta().(*Services).JobService
	for _, rs := range s.RootModule().Resources {
		if _, ok := jobPropertyTypes[rs.Type]; ok {
			_, err := jobService.GetJob(rs.Primary.Attributes["name"])
			// TODO does this really check anything?
			if err == nil {
				return fmt.Errorf("Pipeline property still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
