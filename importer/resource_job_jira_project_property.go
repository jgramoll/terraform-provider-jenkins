package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobJiraProjectProperty"] = jobJiraProjectPropertyCode
}

func jobJiraProjectPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobJiraProjectProperty)
	return fmt.Sprintf(`
  property {
    type = "JiraProjectProperty"
    plugin="%s"
  }
`, property.Plugin)
}
