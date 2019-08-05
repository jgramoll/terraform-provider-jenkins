package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func ensureGitScmBranches(branches *client.GitScmBranches) error {
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

func jobGitScmBranchesCode(branches *client.GitScmBranches) string {
	code := ""
	for i, item := range *branches.Items {
		code += fmt.Sprintf(`
resource "jenkins_job_git_scm_branch" "branch_%v" {
	scm = "${jenkins_job_git_scm.main.id}"

	name = "%v"
}
`, i+1, item.Name)
	}
	return code
}

func jobGitScmBranchesImportScript(jobName string, definitionId string, branches *client.GitScmBranches) string {
	code := ""
	for i, item := range *branches.Items {
		code += fmt.Sprintf(`
terraform import jenkins_job_git_scm_branch.branch_%v "%v"
`, i+1, provider.ResourceJobGitScmBranchId(jobName, definitionId, item.Id))
	}
	return code
}
