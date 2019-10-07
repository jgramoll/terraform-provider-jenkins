package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobPropertyTrackerAction"] = jobDeclarativeJobPropertyTrackerActionCode
}

func jobDeclarativeJobPropertyTrackerActionCode(actionInterface client.JobAction) string {
	action := actionInterface.(*client.JobDeclarativeJobPropertyTrackerAction)
	return fmt.Sprintf(`
resource "jenkins_job_declarative_job_property_tracker_action" "main" {
	job = "${jenkins_job.main.name}"
	plugin = "%v"
}
`, action.Plugin)
}
