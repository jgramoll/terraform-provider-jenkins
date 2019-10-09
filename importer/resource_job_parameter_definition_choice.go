package main

import (
	"fmt"
	"strings"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobParameterDefinitionCodeFuncs["*client.JobParameterDefinitionChoice"] = jobParameterDefinitionChoiceCode
}

func jobParameterDefinitionChoiceCode(
	propertyIndex string, parameterIndex string, propertyInterface client.JobParameterDefinition,
) string {
	definition := propertyInterface.(*client.JobParameterDefinitionChoice)

	choices := *definition.Choices.Items.Items
	for i, c := range choices {
		choices[i] = fmt.Sprintf(`"%v"`, c)
	}
	choicesString := fmt.Sprintf("[%v]", strings.Join(choices, ", "))
	return fmt.Sprintf(`
resource "jenkins_job_parameter_definition_choice" "parameter_%v" {
	property = "${jenkins_job_parameters_definition_property.property_%v.id}"

	name        = "%v"
	description = "%v"
	choices     = %v
}
`, parameterIndex, propertyIndex,
		definition.Name, definition.Description, choicesString)
}
