package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobParametersDefinitionProperty"] = jobParametersDefinitionPropertyCode
}

func jobParametersDefinitionPropertyCode(propertyIndex string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	return fmt.Sprintf(`
resource "jenkins_job_parameters_definition_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex) + jobParameterDefinitionsCode(propertyIndex, property.ParameterDefinitions)
}
