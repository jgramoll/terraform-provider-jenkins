package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobBuildDiscarderProperty"] = jobBuildDiscarderPropertyCode
}

func jobBuildDiscarderPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobBuildDiscarderProperty)
	return `
resource "jenkins_job_build_discarder_property" "main" {
	job = "${jenkins_job.main.id}"
}
` + jobBuildDiscarderPropertyStrategyCode(property.Strategy.Item)
}
