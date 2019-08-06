package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobAction"] = jobDeclarativeJobActionCode
	jobActionImportScriptFuncs["*client.JobDeclarativeJobAction"] = jobDeclarativeJobActionImportScript
}

func jobDeclarativeJobActionCode(actionInterface client.JobAction) string {
	action := actionInterface.(*client.JobDeclarativeJobAction)
	return fmt.Sprintf(`
resource "jenkins_job_declarative_job_action" "main" {
	job = "${jenkins_job.main.name}"
	plugin = "%v"
}
`, action.Plugin)
}

func jobDeclarativeJobActionImportScript(jobName string, action client.JobAction) string {
	return fmt.Sprintf(`
terraform import jenkins_job_declarative_job_action.main "%v"
`, provider.ResourceJobActionId(jobName, action.GetId()))
}
