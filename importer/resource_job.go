package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func ensureJob(job *client.Job) error {
	if job.Id == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		job.Id = id.String()
	}
	if err := ensureJobActions(job.Actions); err != nil {
		return err
	}
	if err := ensureJobProperties(job.Properties); err != nil {
		return err
	}
	return ensureJobDefinition(job.Definition)
}

func jobCode(job *client.Job) string {
	return strings.TrimSpace(fmt.Sprintf(`
resource "jenkins_job" "main" {
	name     = "%v"
	plugin   = "%v"
	disabled = %v
}
`, job.Name, job.Plugin, job.Disabled)+
		jobActionsCode(job.Actions)+
		jobDefinitionCode(job.Definition)+
		jobPropertiesCode(job.Properties)) + "\n"
}

func jobImportScript(job *client.Job) string {
	return strings.TrimSpace(fmt.Sprintf(`
terraform init

terraform import jenkins_job.main "%v"
`, job.Name)+
		jobActionsImportScript(job.Name, job.Actions)+
		jobDefinitionsImportScript(job.Name, job.Definition)+
		jobPropertiesImportScript(job.Name, job.Properties)) + "\n"
}
