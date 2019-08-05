package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDatadogJobProperty"] = jobDatadogJobPropertyCode
}

func jobDatadogJobPropertyCode(client.JobProperty) string {
	return `
resource "jenkins_job_datadog_job_property" "main" {
	job = "${jenkins_job.main.id}"
}
`
}
