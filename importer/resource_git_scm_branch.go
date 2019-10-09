package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

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
