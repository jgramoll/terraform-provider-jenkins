package main

import (
	"fmt"
	"strings"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobCode(job *client.Job) string {
	return strings.TrimSpace(fmt.Sprintf(`
resource "jenkins_job" "main" {
  name     = "%v"
  plugin   = "%v"
  disabled = %v
%s%s%s}
`, job.Name, job.Plugin, job.Disabled,
		jobActionsCode(job.Actions),
		jobDefinitionCode(job.Definition),
		jobPropertiesCode(job.Properties))) + "\n"
}

func jobImportScript(job *client.Job) string {
	return fmt.Sprintf(`
terraform init

terraform import jenkins_job.main "%v"
`, job.Name)
}
