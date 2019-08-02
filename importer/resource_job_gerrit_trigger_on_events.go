package main

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureJobTriggerFuncs["*client.JobGerritTrigger"] = ensureJobGerritTrigger
}

func ensureJobGerritTriggerOnEvents(events *client.JobGerritTriggerOnEvents) error {
	if events == nil || events.Items == nil {
		return nil
	}
	for _, item := range *events.Items {
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
