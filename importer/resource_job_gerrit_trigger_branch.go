package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func ensureJobGerritTriggerBranches(branches *client.JobGerritTriggerBranches) error {
	if branches == nil || branches.Items == nil {
		return nil
	}
	for _, item := range *branches.Items {
		if item.Id == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.Id = id.String()
		}
	}
	return nil
}

func jobGerritTriggerBranchesCode(branches *client.JobGerritTriggerBranches) string {
	code := ""
	for i, item := range *branches.Items {
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_branch" "branch_%v" {
	project = "${jenkins_job_gerrit_project.main.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, i+1, item.CompareType, item.Pattern)
	}
	return code
}
