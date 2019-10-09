package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyCode
}

func jobBuildDiscarderPropertyCode(propertyIndex string, propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return fmt.Sprintf(`
resource "jenkins_job_build_discarder_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex) + jobBuildDiscarderPropertyStrategyCode(propertyIndex, property.Strategy.Item)
}
