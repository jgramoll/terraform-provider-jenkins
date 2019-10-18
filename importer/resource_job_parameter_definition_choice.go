package main

import (
	"fmt"
	"strings"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobParameterDefinitionCodeFuncs["*client.JobParameterDefinitionChoice"] = jobParameterDefinitionChoiceCode
}

func jobParameterDefinitionChoiceCode(propertyInterface client.JobParameterDefinition) string {
	definition := propertyInterface.(*client.JobParameterDefinitionChoice)

	choices := *definition.Choices.Items.Items
	for i, c := range choices {
		choices[i] = fmt.Sprintf(`"%v"`, c)
	}
	choicesString := fmt.Sprintf("[%v]", strings.Join(choices, ", "))
	return fmt.Sprintf(`
    parameter {
      type = "ChoiceParameterDefinition"
      name = "%s"
      description = "%s"
      choices = %s
    }
`, definition.Name, definition.Description, choicesString)
}
