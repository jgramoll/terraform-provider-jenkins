package main

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	ensureJobPropertyFuncs["*client.JobParametersDefinitionProperty"] = ensureJobParametersDefinitionProperty
	jobPropertyCodeFuncs["*client.JobParametersDefinitionProperty"] = jobParametersDefinitionPropertyCode
	jobPropertyImportScriptFuncs["*client.JobParametersDefinitionProperty"] = jobParametersDefinitionPropertyImportScript
}

func ensureJobParametersDefinitionProperty(propertyInterface client.JobProperty) error {
	property := propertyInterface.(*client.JobParametersDefinitionProperty)

	for _, item := range *property.ParameterDefinitions.Items {
		if item.GetId() == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.SetId(id.String())
		}
	}
	return nil
}

func jobParametersDefinitionPropertyCode(propertyIndex string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	return fmt.Sprintf(`
resource "jenkins_job_parameters_definition_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex) + jobParameterDefinitionsCode(propertyIndex, property.ParameterDefinitions)
}

func jobParametersDefinitionPropertyImportScript(
	propertyIndex string, jobName string, propertyInterface client.JobProperty,
) string {
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	propertyId := propertyInterface.GetId()
	return fmt.Sprintf(`
terraform import jenkins_job_parameters_definition_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, propertyId)) +
		jobParameterDefinitionsImportScript(propertyIndex, jobName, propertyId, property.ParameterDefinitions)
}
