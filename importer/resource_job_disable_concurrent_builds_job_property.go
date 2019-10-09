package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDisableConcurrentBuildsJobProperty"] = jobDisableConcurrentBuildsJobPropertyCode
}

func jobDisableConcurrentBuildsJobPropertyCode(propertyIndex string, property client.JobProperty) string {
	return fmt.Sprintf(`
resource "jenkins_job_disable_concurrent_builds_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex)
}
