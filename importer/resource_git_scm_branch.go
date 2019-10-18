package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGitScmBranchesCode(branches *client.GitScmBranches) string {
	code := ""
	for _, item := range *branches.Items {
		code += fmt.Sprintf(`
      branch {
        name = "%s"
      }
`, item.Name)
	}
	return code
}
