package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDisableConcurrentBuildsJobProperty"] = jobDisableConcurrentBuildsJobPropertyCode
	jobPropertyImportScriptFuncs["*client.JobDisableConcurrentBuildsJobProperty"] = jobDisableConcurrentBuildsJobPropertyImportScript
}

func jobDisableConcurrentBuildsJobPropertyCode(propertyIndex string, property client.JobProperty) string {
	return fmt.Sprintf(`
resource "jenkins_job_disable_concurrent_builds_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex)
}

func jobDisableConcurrentBuildsJobPropertyImportScript(propertyIndex string, jobName string, p client.JobProperty) string {
	return fmt.Sprintf(`
terraform import jenkins_job_disable_concurrent_builds_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, p.GetId()))
}
