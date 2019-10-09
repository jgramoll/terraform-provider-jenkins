package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerEventInitFunc[client.PluginDraftPublishedEventType] = func() jobGerritTriggerEvent {
		return newJobGerritTriggerDraftPublishedEvent()
	}
}

type jobGerritTriggerDraftPublishedEvent struct {
	Type client.JobGerritTriggerOnEventType `mapstructure:"type"`
}

func newJobGerritTriggerDraftPublishedEvent() *jobGerritTriggerDraftPublishedEvent {
	return &jobGerritTriggerDraftPublishedEvent{
		Type: client.PluginDraftPublishedEventType,
	}
}

func (event *jobGerritTriggerDraftPublishedEvent) fromClientTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
	_, ok := clientEventInterface.(*client.JobGerritTriggerPluginDraftPublishedEvent)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobGerritTriggerPluginDraftPublishedEvent, got %s",
			reflect.TypeOf(clientEventInterface).String())
	}
	return newJobGerritTriggerDraftPublishedEvent(), nil
}

func (event *jobGerritTriggerDraftPublishedEvent) toClientTriggerEvent() (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	return clientEvent, nil
}
