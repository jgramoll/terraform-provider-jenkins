package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritTriggerPatchsetCreatedEvent struct {
	ExcludeDrafts        bool `mapstructure:"exclude_drafts"`
	ExcludeTrivialRebase bool `mapstructure:"exclude_trivial_rebase"`
	ExcludeNoCodeChange  bool `mapstructure:"exclude_no_code_change"`
	ExcludePrivateState  bool `mapstructure:"exclude_private_state"`
	ExcludeWipState      bool `mapstructure:"exclude_wip_state"`
}

func newJobGerritTriggerPatchsetCreatedEvent() *jobGerritTriggerPatchsetCreatedEvent {
	return &jobGerritTriggerPatchsetCreatedEvent{}
}

func (e *jobGerritTriggerPatchsetCreatedEvent) fromClientJobTriggerEvent(clientEventInterface client.JobGerritTriggerOnEvent) (jobGerritTriggerEvent, error) {
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

func (event *jobGerritTriggerPatchsetCreatedEvent) toClientJobTriggerEvent(eventId string) (client.JobGerritTriggerOnEvent, error) {
	clientEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	clientEvent.Id = eventId
	clientEvent.ExcludeDrafts = event.ExcludeDrafts
	clientEvent.ExcludeTrivialRebase = event.ExcludeTrivialRebase
	clientEvent.ExcludeNoCodeChange = event.ExcludeNoCodeChange
	clientEvent.ExcludePrivateState = event.ExcludePrivateState
	clientEvent.ExcludeWipState = event.ExcludeWipState
	return clientEvent, nil
}

func (event *jobGerritTriggerPatchsetCreatedEvent) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("exclude_drafts", event.ExcludeDrafts); err != nil {
		return err
	}
	if err := d.Set("exclude_trivial_rebase", event.ExcludeTrivialRebase); err != nil {
		return err
	}
	if err := d.Set("exclude_no_code_change", event.ExcludeNoCodeChange); err != nil {
		return err
	}
	if err := d.Set("exclude_private_state", event.ExcludePrivateState); err != nil {
		return err
	}
	return d.Set("exclude_wip_state", event.ExcludeWipState)
}
