package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerEventInitFunc[client.PluginPatchsetCreatedEventType] = func() jobGerritTriggerEvent {
		return newJobGerritTriggerPatchsetCreatedEvent()
	}
}

type jobGerritTriggerPatchsetCreatedEvent struct {
	Type client.JobGerritTriggerOnEventType `mapstructure:"type"`

	ExcludeDrafts        bool `mapstructure:"exclude_drafts"`
	ExcludeTrivialRebase bool `mapstructure:"exclude_trivial_rebase"`
	ExcludeNoCodeChange  bool `mapstructure:"exclude_no_code_change"`
	ExcludePrivateState  bool `mapstructure:"exclude_private_state"`
	ExcludeWipState      bool `mapstructure:"exclude_wip_state"`
}

func newJobGerritTriggerPatchsetCreatedEvent() *jobGerritTriggerPatchsetCreatedEvent {
	return &jobGerritTriggerPatchsetCreatedEvent{
		Type: client.PluginPatchsetCreatedEventType,
	}
}

func (e *jobGerritTriggerPatchsetCreatedEvent) fromClientTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
	clientEvent, ok := clientEventInterface.(*client.JobGerritTriggerPluginPatchsetCreatedEvent)
	if !ok {
		return nil, fmt.Errorf("Unexpected event type got %s, expected *client.JobGerritTriggerPluginPatchsetCreatedEvent",
			reflect.TypeOf(clientEventInterface).String())
	}
	event := newJobGerritTriggerPatchsetCreatedEvent()
	event.ExcludeDrafts = clientEvent.ExcludeDrafts
	event.ExcludeTrivialRebase = clientEvent.ExcludeTrivialRebase
	event.ExcludeNoCodeChange = clientEvent.ExcludeNoCodeChange
	event.ExcludePrivateState = clientEvent.ExcludePrivateState
	event.ExcludeWipState = clientEvent.ExcludeWipState
	return event, nil
}

func (event *jobGerritTriggerPatchsetCreatedEvent) toClientTriggerEvent() (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	clientEvent.ExcludeDrafts = event.ExcludeDrafts
	clientEvent.ExcludeTrivialRebase = event.ExcludeTrivialRebase
	clientEvent.ExcludeNoCodeChange = event.ExcludeNoCodeChange
	clientEvent.ExcludePrivateState = event.ExcludePrivateState
	clientEvent.ExcludeWipState = event.ExcludeWipState
	return clientEvent, nil
}
