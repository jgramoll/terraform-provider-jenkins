package main

import (
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
