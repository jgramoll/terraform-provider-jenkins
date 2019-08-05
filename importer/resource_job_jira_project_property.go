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

func jobJiraProjectPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobJiraProjectProperty)
	return fmt.Sprintf(`
resource "jenkins_job_jira_project_property" "main" {
	job = "${jenkins_job.main.name}"

	plugin = "%v"
}
`, property.Plugin)
}

func jobJiraProjectPropertyImportScript(jobName string, p client.JobProperty) string {
	return fmt.Sprintf(`
terraform import jenkins_job_jira_project_property.main "%v"
`, provider.ResourceJobPropertyId(jobName, p.GetId()))
}
