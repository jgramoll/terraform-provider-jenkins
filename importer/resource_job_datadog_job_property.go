package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDatadogJobProperty"] = jobDatadogJobPropertyCode
	jobPropertyImportScriptFuncs["*client.JobDatadogJobProperty"] = jobDatadogJobPropertyImportScript
}

func jobDatadogJobPropertyCode(client.JobProperty) string {
	return `
resource "jenkins_job_datadog_job_property" "main" {
	job = "${jenkins_job.main.name}"
}
`
}

func jobDatadogJobPropertyImportScript(jobName string, p client.JobProperty) string {
	return fmt.Sprintf(`
terraform import jenkins_job_datadog_job_property.main "%v"
`, provider.ResourceJobPropertyId(jobName, p.GetId()))
}
