package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func ensureJobGerritTriggerFilePaths(filePaths *client.JobGerritTriggerFilePaths) error {
	if filePaths == nil || filePaths.Items == nil {
		return nil
	}
	for _, item := range *filePaths.Items {
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

func jobGerritTriggerFilePathsCode(projectIndex int, filePaths *client.JobGerritTriggerFilePaths) string {
	code := ""
	for i, item := range *filePaths.Items {
		code += fmt.Sprintf(`
resource "jenkins_job_gerrit_file_path" "file_path_%v_%v" {
	project = "${jenkins_job_gerrit_project.project_%v.id}"

	compare_type = "%v"
	pattern      = "%v"
}
`, i+1, projectIndex, projectIndex, item.CompareType, item.Pattern)
	}
	return code
}

func jobGerritTriggerFilePathsImportScript(
	projectIndex int, jobName string, propertyId string, triggerId string,
	projectId string, filePaths *client.JobGerritTriggerFilePaths,
) string {
	code := ""
	for i, item := range *filePaths.Items {
		id := provider.ResourceJobGerritFilePathId(
			jobName, propertyId, triggerId, projectId, item.Id)
		code += fmt.Sprintf(`
terraform import jenkins_job_gerrit_file_path.file_path_%v_%v "%v"
`, i+1, projectIndex, id)
	}
	return code
}
