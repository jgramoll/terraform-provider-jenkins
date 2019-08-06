package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerChangeMergedEvent struct {
}

func newJobGerritTriggerChangeMergedEvent() *jobGerritTriggerChangeMergedEvent {
	return &jobGerritTriggerChangeMergedEvent{}
}

func (e *jobGerritTriggerChangeMergedEvent) fromClientJobTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
	_, ok := clientEventInterface.(*client.JobGerritTriggerPluginChangeMergedEvent)
	if !ok {
		return nil, fmt.Errorf("Unexpected event type got %s, expected *client.JobGerritTriggerPluginChangeMergedEvent",
			reflect.TypeOf(clientEventInterface).String())
	}
	event := newJobGerritTriggerChangeMergedEvent()
	return event, nil
}

func (event *jobGerritTriggerChangeMergedEvent) toClientJobTriggerEvent(eventId string) (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginChangeMergedEvent()
	clientEvent.Id = eventId
	return clientEvent, nil
}

func (event *jobGerritTriggerChangeMergedEvent) setResourceData(d *schema.ResourceData) error {
	return nil
}
