package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerProjectsCodeFunc func(*client.JobGerritTriggerProject) string

var jobGerritTriggerProjectsCodeFuncs = map[string]jobGerritTriggerProjectsCodeFunc{}

func jobGerritTriggerProjectsCode(
	propertyIndex string, triggerIndex string, projects *client.JobGerritTriggerProjects,
) string {
	code := ""
	for i, item := range *projects.Items {
		projectIndex := fmt.Sprintf("%v_%v", triggerIndex, i+1)
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_project" "project_%v" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_%v.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, projectIndex, triggerIndex, item.CompareType, item.Pattern) +
			jobGerritTriggerBranchesCode(projectIndex, item.Branches) +
			jobGerritTriggerFilePathsCode(projectIndex, item.FilePaths)
	}
	return code
}
