package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerDraftPublishedEvent struct {
}

func newJobGerritTriggerDraftPublishedEvent() *jobGerritTriggerDraftPublishedEvent {
	return &jobGerritTriggerDraftPublishedEvent{}
}

func (event *jobGerritTriggerDraftPublishedEvent) fromClientJobTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
	_, ok := clientEventInterface.(*client.JobGerritTriggerPluginDraftPublishedEvent)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobGerritTriggerPluginDraftPublishedEvent, got %s",
			reflect.TypeOf(clientEventInterface).String())
	}
	return newJobGerritTriggerDraftPublishedEvent(), nil
}

func (event *jobGerritTriggerDraftPublishedEvent) toClientJobTriggerEvent(eventId string) (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	clientEvent.Id = eventId
	return clientEvent, nil
}

func (event *jobGerritTriggerDraftPublishedEvent) setResourceData(d *schema.ResourceData) error {
	return nil
}
