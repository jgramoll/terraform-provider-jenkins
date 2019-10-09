package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDisableConcurrentBuildsJobProperty"] = jobDisableConcurrentBuildsJobPropertyCode
}

func jobDisableConcurrentBuildsJobPropertyCode(property client.JobProperty) string {
	return `
  property {
    type = "DisableConcurrentBuildsJobProperty"
  }
`
}
