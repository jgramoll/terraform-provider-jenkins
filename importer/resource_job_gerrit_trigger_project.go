package main

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func ensureJobGerritTriggerProjects(projects *client.JobGerritTriggerProjects) error {
	if projects == nil || projects.Items == nil {
		return nil
	}
	for _, item := range *projects.Items {
		if item.Id == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.Id = id.String()
		}
		if err := ensureJobGerritTriggerBranches(item.Branches); err != nil {
			return err
		}
	}
	return nil
}
