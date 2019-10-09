package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGerritTriggerFilePathsCode(projectIndex string, filePaths *client.JobGerritTriggerFilePaths) string {
	code := ""
	if filePaths == nil || filePaths.Items == nil {
		return code
	}
	for i, item := range *filePaths.Items {
		filePathIndex := fmt.Sprintf("%v_%v", projectIndex, i+1)
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_file_path" "file_path_%v" {
	project = "${jenkins_job_gerrit_project.project_%v.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, filePathIndex, projectIndex, item.CompareType, item.Pattern)
	}
	return code
}
