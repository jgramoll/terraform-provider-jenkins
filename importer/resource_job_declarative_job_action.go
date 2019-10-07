package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobActionCodeFuncs["*client.JobDeclarativeJobAction"] = jobDeclarativeJobActionCode
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
