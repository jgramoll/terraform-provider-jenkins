package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyCode
}

func jobBuildDiscarderPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return fmt.Sprintf(`
  property {
    type = "BuildDiscarderProperty"
%s  }
`, jobBuildDiscarderPropertyStrategyCode(property.Strategy.Item))
}
