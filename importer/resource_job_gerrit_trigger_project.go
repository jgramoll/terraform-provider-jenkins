package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
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
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_project" "project_%v" {
	trigger = "${jenkins_job_gerrit_trigger.main.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, i+1, item.CompareType, item.Pattern) + jobGerritTriggerBranchesCode(item.Branches)
	}
	return code
}
