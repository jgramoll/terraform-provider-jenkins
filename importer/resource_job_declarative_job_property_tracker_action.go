package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobPropertyTrackerAction"] = jobDeclarativeJobPropertyTrackerActionCode
	jobActionImportScriptFuncs["*client.JobDeclarativeJobPropertyTrackerAction"] = jobDeclarativeJobPropertyTrackerActionImportScript
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

func jobDeclarativeJobPropertyTrackerActionImportScript(jobName string, action client.JobAction) string {
	return fmt.Sprintf(`
terraform import jenkins_job_declarative_job_property_tracker_action.main "%v"
`, provider.ResourceJobActionId(jobName, action.GetId()))
}
