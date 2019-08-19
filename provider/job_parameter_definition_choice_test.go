package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func TestAccJobParameterDefinitionChoiceBasic(t *testing.T) {
	var jobRef client.Job
	var parameters []client.JobParameterDefinition
	jobName := fmt.Sprintf("%s/tf-acc-test-%s", jenkinsFolder, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jobResourceName := "jenkins_job.main"
	parameter1 := "jenkins_job_parameter_definition_choice.param_1"
	parameter2 := "jenkins_job_parameter_definition_choice.param_2"
	parameterName := "paramName"
	newParameterName := "newParamName"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, parameterName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", parameterName+"_1"),
					resource.TestCheckResourceAttr(parameter1, "description", "choice_description"),
					resource.TestCheckResourceAttr(parameter1, "choices.0", "1"),
					resource.TestCheckResourceAttr(parameter1, "choices.1", "3"),
					resource.TestCheckResourceAttr(parameter1, "choices.2", "4"),
					resource.TestCheckResourceAttr(parameter2, "name", parameterName+"_2"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobParameterDefintion(&jobRef, []string{
						parameter1,
						parameter2,
					}, &parameters, ensureJobParameterDefinitionChoice),
				),
			},
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, newParameterName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", newParameterName+"_1"),
					resource.TestCheckResourceAttr(parameter1, "description", "choice_description"),
					resource.TestCheckResourceAttr(parameter1, "choices.0", "1"),
					resource.TestCheckResourceAttr(parameter1, "choices.1", "3"),
					resource.TestCheckResourceAttr(parameter1, "choices.2", "4"),
					resource.TestCheckResourceAttr(parameter2, "name", newParameterName+"_2"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobParameterDefintion(&jobRef, []string{
						parameter1,
						parameter2,
					}, &parameters, ensureJobParameterDefinitionChoice),
				),
			},
			{
				ResourceName:  parameter1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Invalid parameter id"),
			},
			{
				ResourceName: parameter1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(parameters) == 0 {
						return "", fmt.Errorf("no parameters to import")
					}
					propertyId := (*jobRef.Properties.Items)[0].GetId()
					return ResourceJobParameterDefinitionId(jobName, propertyId, parameters[0].GetId()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: parameter2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(parameters) == 0 {
						return "", fmt.Errorf("no parameters to import")
					}
					propertyId := (*jobRef.Properties.Items)[0].GetId()
					return ResourceJobParameterDefinitionId(jobName, propertyId, parameters[1].GetId()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, parameterName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", parameterName+"_1"),
					resource.TestCheckResourceAttr(parameter1, "description", "choice_description"),
					resource.TestCheckResourceAttr(parameter1, "choices.0", "1"),
					resource.TestCheckResourceAttr(parameter1, "choices.1", "3"),
					resource.TestCheckResourceAttr(parameter1, "choices.2", "4"),
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobParameterDefintion(&jobRef, []string{
						parameter1,
					}, &parameters, ensureJobParameterDefinitionChoice),
				),
			},
			{
				Config: testAccJobParameterDefinitionChoiceConfigBasic(jobName, parameterName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckJobExists(jobResourceName, &jobRef),
					testAccCheckJobParameterDefintion(&jobRef, []string{}, &parameters, ensureJobParameterDefinitionChoice),
				),
			},
		},
	})
}

func testAccJobParameterDefinitionChoiceConfigBasic(jobName string, paramName string, count int) string {
	definitions := ""
	for i := 1; i <= count; i++ {
		definitions += fmt.Sprintf(`
resource "jenkins_job_parameter_definition_choice" "param_%v" {
	property = "${jenkins_job_parameters_definition_property.prop_1.id}"

	name = "%v_%v"
	description = "choice_description"
	choices = [
		"1",
		"3",
		"4"
	]
}`, i, paramName, i)
	}

	return testAccJobParametersDefinitionPropertyConfigBasic(jobName, 1) + definitions
}

func ensureJobParameterDefinitionChoice(definitionInterface client.JobParameterDefinition, resource *terraform.ResourceState) error {
	_, err := newJobParameterDefinitionChoice().fromClientJobParameterDefintion(definitionInterface)
	if err != nil {
		return err
	}
	definition := definitionInterface.(*client.JobParameterDefinitionChoice)
	if definition.Name != resource.Primary.Attributes["name"] {
		return fmt.Errorf("Expected definition name to be %v, got %v", resource.Primary.Attributes["name"], definition.Name)
	}
	if definition.Description != resource.Primary.Attributes["description"] {
		return fmt.Errorf("Expected definition description to be %v, got %v", resource.Primary.Attributes["description"], definition.Description)
	}
	for i, choice := range *definition.Choices.Items.Items {
		resourceChoice := resource.Primary.Attributes[fmt.Sprintf("choices.%v", i)]
		if resourceChoice != choice {
			return fmt.Errorf("Expected definition choices to be %v, got %v", resourceChoice, choice)
		}
	}
	return nil
}
