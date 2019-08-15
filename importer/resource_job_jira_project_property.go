package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobPropertyCodeFuncs["*client.JobJiraProjectProperty"] = jobJiraProjectPropertyCode
	jobPropertyImportScriptFuncs["*client.JobJiraProjectProperty"] = jobJiraProjectPropertyImportScript
}

func jobJiraProjectPropertyCode(propertyIndex int, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobJiraProjectProperty)
	return fmt.Sprintf(`
resource "jenkins_job_jira_project_property" "property_%v" {
	job = "${jenkins_job.main.name}"

	plugin = "%v"
}
`, propertyIndex, property.Plugin)
}

func jobJiraProjectPropertyImportScript(propertyIndex int, jobName string, p client.JobProperty) string {
	return fmt.Sprintf(`
terraform import jenkins_job_jira_project_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, p.GetId()))
}
