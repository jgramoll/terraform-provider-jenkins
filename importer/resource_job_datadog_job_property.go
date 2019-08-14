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

func jobDatadogJobPropertyCode(propertyIndex int, property client.JobProperty) string {
	return fmt.Sprintf(`
resource "jenkins_job_datadog_job_property" "property_%v" {
	job = "${jenkins_job.main.name}"
}
`, propertyIndex)
}

func jobDatadogJobPropertyImportScript(propertyIndex int, jobName string, p client.JobProperty) string {
	return fmt.Sprintf(`
terraform import jenkins_job_datadog_job_property.property_%v "%v"
`, propertyIndex, provider.ResourceJobPropertyId(jobName, p.GetId()))
}
