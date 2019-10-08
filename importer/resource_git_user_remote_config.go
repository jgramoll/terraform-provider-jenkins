package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func ensureGitUserRemoteConfigs(configs *client.GitUserRemoteConfigs) error {
	if configs == nil || configs.Items == nil {
		return nil
	}
	for _, item := range *configs.Items {
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

func jobGitScmUserRemoteConfigsCode(configs *client.GitUserRemoteConfigs) string {
	code := ""
	for i, item := range *configs.Items {
		// Need $ -> $$ for tf escape
		refspec := strings.ReplaceAll(item.Refspec, "$", "$$")
		code += fmt.Sprintf(`
resource "jenkins_job_git_scm_user_remote_config" "config_%v" {
	scm = "${jenkins_job_git_scm.main.id}"

	refspec        = "%v"
	url            = "%v"
	credentials_id = "%v"
}
`, i+1, refspec, item.Url, item.CredentialsId)
	}
	return code
}
