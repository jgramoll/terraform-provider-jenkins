package main

import "github.com/jgramoll/terraform-provider-jenkins/client"

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobAction"] = jobDeclarativeJobActionCode
}

func jobDeclarativeJobActionCode(client.JobAction) string {
	return `
resource "jenkins_job_declarative_job_action" "main" {
	job = "${jenkins_job.main.id}"
}
`
}
