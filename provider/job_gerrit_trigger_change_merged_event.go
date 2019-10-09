package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerEventInitFunc[client.PluginChangeMergedEventType] = func() jobGerritTriggerEvent {
		return newJobGerritTriggerChangeMergedEvent()
	}
}

type jobGerritTriggerChangeMergedEvent struct {
	Type client.JobGerritTriggerOnEventType `mapstructure:"type"`
}

func newJobGerritTriggerChangeMergedEvent() *jobGerritTriggerChangeMergedEvent {
	return &jobGerritTriggerChangeMergedEvent{
		Type: client.PluginChangeMergedEventType,
	}
}

func (e *jobGerritTriggerChangeMergedEvent) fromClientTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
	_, ok := clientEventInterface.(*client.JobGerritTriggerPluginChangeMergedEvent)
	if !ok {
		return nil, fmt.Errorf("Unexpected event type got %s, expected *client.JobGerritTriggerPluginChangeMergedEvent",
			reflect.TypeOf(clientEventInterface).String())
	}
	event := newJobGerritTriggerChangeMergedEvent()
	return event, nil
}

func (event *jobGerritTriggerChangeMergedEvent) toClientTriggerEvent() (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginChangeMergedEvent()
	return clientEvent, nil
}
