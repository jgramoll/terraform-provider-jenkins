package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGerritTriggerBranchesCode(branches *client.JobGerritTriggerBranches) string {
	code := ""
	for _, item := range *branches.Items {
		code += fmt.Sprintf(`
        branch {
          compare_type = "%s"
          pattern      = "%s"
        }
`, item.CompareType, item.Pattern)
	}
	return code
}
