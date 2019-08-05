package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobJiraProjectProperty"] = jobJiraProjectPropertyCode
}

func jobJiraProjectPropertyCode(client.JobProperty) string {
	return `
resource "jenkins_job_jira_project_property" "main" {
	job = "${jenkins_job.main.id}"
}
`
}
