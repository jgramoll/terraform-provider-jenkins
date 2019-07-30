package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerDraftPublishedEvent struct {
}

func newJobGerritTriggerDraftPublishedEvent() *jobGerritTriggerDraftPublishedEvent {
	return &jobGerritTriggerDraftPublishedEvent{}
}

func (event *jobGerritTriggerDraftPublishedEvent) fromClientJobTriggerEvent(clientEvent client.JobGerritTriggerOnEvent) jobGerritTriggerEvent {
	return newJobGerritTriggerDraftPublishedEvent()
}

func (event *jobGerritTriggerDraftPublishedEvent) toClientJobTriggerEvent(eventId string) (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	clientEvent.Id = eventId
	return clientEvent, nil
}

func (event *jobGerritTriggerDraftPublishedEvent) setResourceData(d *schema.ResourceData) error {
	return nil
}
