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
  action {
    type = "DeclarativeJobAction"
    plugin = "%s"
  }
`, action.Plugin)
}
