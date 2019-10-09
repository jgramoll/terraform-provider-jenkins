package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func jobGerritTriggerFilePathsCode(filePaths *client.JobGerritTriggerFilePaths) string {
	code := ""
	if filePaths == nil || filePaths.Items == nil {
		return code
	}
	for _, item := range *filePaths.Items {
		code += fmt.Sprintf(`
        file_path {
          compare_type = "%s"
          pattern      = "%s"
        }
`, item.CompareType, item.Pattern)
	}
	return code
}
