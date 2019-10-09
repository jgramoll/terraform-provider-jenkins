package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobJiraProjectProperty"] = jobJiraProjectPropertyCode
}

func jobJiraProjectPropertyCode(propertyIndex string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobJiraProjectProperty)
	return fmt.Sprintf(`
resource "jenkins_job_jira_project_property" "property_%v" {
	job = "${jenkins_job.main.name}"

	plugin = "%v"
}
`, propertyIndex, property.Plugin)
}
