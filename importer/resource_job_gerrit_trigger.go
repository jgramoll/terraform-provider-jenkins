package main

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	ensureJobTriggerFuncs["*client.JobGerritTrigger"] = ensureJobGerritTrigger
}

func ensureJobGerritTrigger(triggerInterface client.JobTrigger) error {
	trigger := triggerInterface.(*client.JobGerritTrigger)
	if trigger.Id == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		trigger.Id = id.String()
	}
	if err := ensureJobDynamicGerritProjects(trigger.DynamicGerritProjects); err != nil {
		return err
	}
	if err := ensureJobGerritTriggerProjects(trigger.Projects); err != nil {
		return err
	}
	return ensureJobGerritTriggerOnEvents(trigger.TriggerOnEvents)
}
