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

func jobBuildDiscarderPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return `
resource "jenkins_job_build_discarder_property" "main" {
	job = "${jenkins_job.main.name}"
}
` + jobBuildDiscarderPropertyStrategyCode(property.Strategy.Item)
}

func jobBuildDiscarderPropertyImportScript(jobName string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return fmt.Sprintf(`
terraform import jenkins_job_build_discarder_property.main "%v"
`, provider.ResourceJobPropertyId(jobName, property.Id)) +
		jobBuildDiscarderPropertyStrategyImportScript(jobName, property.Id, property.Strategy.Item)
}
