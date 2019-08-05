package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type gitUserRemoteConfigCodeFunc func(client.GitUserRemoteConfig) string

var gitUserRemoteConfigCodeFuncs = map[string]gitUserRemoteConfigCodeFunc{}

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
		code += fmt.Sprintf(`
resource "jenkins_job_git_scm_user_remote_config" "config_%v" {
	scm = "${jenkins_job_git_scm.main.id}"

	refspec        = "%v"
	url            = "%v"
	credentials_id = "%v"
}
`, i+1, item.Refspec, item.Url, item.CredentialsId)
	}
	return code
}
