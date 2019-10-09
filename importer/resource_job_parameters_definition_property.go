package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobParametersDefinitionProperty"] = jobParametersDefinitionPropertyCode
}

func jobParametersDefinitionPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	return fmt.Sprintf(`
  property {
    type = "ParametersDefinitionProperty"
%s  }
`, jobParameterDefinitionsCode(property.ParameterDefinitions))
}
