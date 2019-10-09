package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerProjectsCodeFunc func(*client.JobGerritTriggerProject) string

var jobGerritTriggerProjectsCodeFuncs = map[string]jobGerritTriggerProjectsCodeFunc{}

func jobGerritTriggerProjectsCode(projects *client.JobGerritTriggerProjects) string {
	code := ""
	for _, item := range *projects.Items {
		code += fmt.Sprintf(`
      gerrit_project {
        compare_type = "%s"
        pattern      = "%s"
%s%s      }
`, item.CompareType, item.Pattern,
			jobGerritTriggerBranchesCode(item.Branches),
			jobGerritTriggerFilePathsCode(item.FilePaths))
	}
	return code
}
