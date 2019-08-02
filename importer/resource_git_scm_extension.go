package main

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func ensureGitScmExtensions(extensions *client.GitScmExtensions) error {
	if extensions == nil || extensions.Items == nil {
		return nil
	}
	for _, item := range *extensions.Items {
		if item.GetId() == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.SetId(id.String())
		}
	}
	return nil
}
