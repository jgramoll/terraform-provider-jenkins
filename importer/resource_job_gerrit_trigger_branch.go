package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGerritTriggerBranchesCode(projectIndex string, branches *client.JobGerritTriggerBranches) string {
	code := ""
	for i, item := range *branches.Items {
		branchIndex := fmt.Sprintf("%v_%v", projectIndex, i+1)
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_branch" "branch_%v" {
	project = "${jenkins_job_gerrit_project.project_%v.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, branchIndex, projectIndex, item.CompareType, item.Pattern)
	}
	return code
}
