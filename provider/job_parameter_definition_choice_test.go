package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobParameterDefinitionChoiceBasic(t *testing.T) {
	var jobRef client.Job
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	parameterName := "paramName"
	newParameterName := "newParamName"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, parameterName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.type", "ChoiceParameterDefinition"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.name", parameterName),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.description", "which env to target"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.0", "1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.1", "3"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.2", "4"),
				),
			},
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, newParameterName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.type", "ChoiceParameterDefinition"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.name", newParameterName),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.description", "which env to target"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.0", "1"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.1", "3"),
					resource.TestCheckResourceAttr(jobResourceName, "property.0.parameter.0.choices.2", "4"),
				),
			},
		},
	})
}

func testAccJobParameterDefinitionChoiceConfigBasic(jobName string, parameterName string) string {
	return fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%s"
	plugin   = "workflow-job@2.33"

	property {
		type = "ParametersDefinitionProperty"

		parameter {
			type = "ChoiceParameterDefinition"
			name = "%s"
			description = "which env to target"
			choices = ["1", "3", "4"]
		}
	}
}`, jobName, parameterName)
}
