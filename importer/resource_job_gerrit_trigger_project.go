package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

type jobGerritTriggerProjectsCodeFunc func(*client.JobGerritTriggerProject) string

var jobGerritTriggerProjectsCodeFuncs = map[string]jobGerritTriggerProjectsCodeFunc{}

func ensureJobGerritTriggerProjects(projects *client.JobGerritTriggerProjects) error {
	if projects == nil || projects.Items == nil {
		return nil
	}
	for _, item := range *projects.Items {
		if item.Id == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.Id = id.String()
		}
		if err := ensureJobGerritTriggerBranches(item.Branches); err != nil {
			return err
		}
	}
	return nil
}

func jobGerritTriggerProjectsCode(projects *client.JobGerritTriggerProjects) string {
	code := ""
	for i, item := range *projects.Items {
		projectIndex := i + 1
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_project" "project_%v" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, projectIndex, item.CompareType, item.Pattern) + jobGerritTriggerBranchesCode(projectIndex, item.Branches)
	}
	return code
}

func jobGerritTriggerProjectsImportScript(
	jobName string, propertyId string, triggerId string,
	projects *client.JobGerritTriggerProjects,
) string {
	code := ""
	for i, item := range *projects.Items {
		projectIndex := i + 1
		code += fmt.Sprintf(`
terraform import jenkins_job_gerrit_project.project_%v "%v"
`, projectIndex, provider.ResourceJobGerritProjectId(jobName, propertyId, triggerId, item.Id)) +
			jobGerritTriggerBranchesImportScript(
				projectIndex, jobName, propertyId, triggerId, item.Id, item.Branches)
	}
	return code
}
