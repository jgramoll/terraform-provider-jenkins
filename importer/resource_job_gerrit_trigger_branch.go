package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
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

func jobGerritTriggerBranchesCode(projectIndex int, branches *client.JobGerritTriggerBranches) string {
	code := ""
	for i, item := range *branches.Items {
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_branch" "branch_%v_%v" {
	project = "${jenkins_job_gerrit_project.project_%v.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, projectIndex, i+1, projectIndex, item.CompareType, item.Pattern)
	}
	return code
}

func jobGerritTriggerBranchesImportScript(
	projectIndex int, jobName string, propertyId string, triggerId string,
	projectId string, branches *client.JobGerritTriggerBranches,
) string {
	code := ""
	for i, item := range *branches.Items {
		id := provider.ResourceJobGerritBranchId(
			jobName, propertyId, triggerId, projectId, item.Id)
		code += fmt.Sprintf(`
terraform import jenkins_job_gerrit_branch.branch_%v_%v "%v"
`, i+1, projectIndex, id)
	}
	return code
}
