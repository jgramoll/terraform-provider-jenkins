package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDatadogJobProperty"] = jobDatadogJobPropertyCode
}

func jobDatadogJobPropertyCode(propertyIndex string, property client.JobProperty) string {
	return fmt.Sprintf(`
resource "jenkins_job_datadog_job_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex)
}
