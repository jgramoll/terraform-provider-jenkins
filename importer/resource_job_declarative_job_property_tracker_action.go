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
  action {
    type = "DeclarativeJobPropertyTrackerAction"
    plugin = "%s"
  }
`, action.Plugin)
}
