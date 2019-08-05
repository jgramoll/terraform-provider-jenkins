package main

import "github.com/jgramoll/terraform-provider-jenkins/client"

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobPropertyTrackerAction"] = jobDeclarativeJobPropertyTrackerActionCode
}

func jobDeclarativeJobPropertyTrackerActionCode(client.JobAction) string {
	return `
resource "jenkins_job_declarative_job_property_tracker_action" "main" {
	job = "${jenkins_job.main.id}"
}
`
}
