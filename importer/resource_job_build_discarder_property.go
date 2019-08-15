package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobPropertyCodeFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyCode
	jobPropertyImportScriptFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyImportScript
}

func jobBuildDiscarderPropertyCode(propertyIndex int, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return fmt.Sprintf(`
resource "jenkins_job_build_discarder_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex) + jobBuildDiscarderPropertyStrategyCode(propertyIndex, property.Strategy.Item)
}

func jobBuildDiscarderPropertyImportScript(propertyIndex int, jobName string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return fmt.Sprintf(`
terraform import jenkins_job_build_discarder_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, property.Id)) +
		jobBuildDiscarderPropertyStrategyImportScript(jobName, property.Id, property.Strategy.Item)
}
