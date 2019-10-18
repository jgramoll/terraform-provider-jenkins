package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type interfaceJobTriggerEvents []map[string]interface{}

type jobTriggerEventInit func() jobGerritTriggerEvent

var jobTriggerEventInitFunc = map[client.JobGerritTriggerOnEventType]jobTriggerEventInit{}

func (triggerEvents *interfaceJobTriggerEvents) toClientTriggerEvents() (*client.JobGerritTriggerOnEvents, error) {
	clientTriggerEvents := client.NewJobGerritTriggerOnEvents()
	for _, mapData := range *triggerEvents {
		triggerEventTypeString, ok := mapData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize job gerrit trigger event, missing type")
		}
		triggerEventType, err := client.ParseJobGerritTriggerOnEventType(triggerEventTypeString)
		if err != nil {
			return nil, err
		}
		initFunc := jobTriggerEventInitFunc[triggerEventType]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job gerrit trigger event, missing init func")
		}
		triggerEvent := initFunc()
		if err := mapstructure.Decode(mapData, &triggerEvent); err != nil {
			return nil, err
		}
		clientTriggerEvent, err := triggerEvent.toClientTriggerEvent()
		if err != nil {
			return nil, err
		}
		clientTriggerEvents = clientTriggerEvents.Append(clientTriggerEvent)
	}
	return clientTriggerEvents, nil
}

func (*interfaceJobTriggerEvents) fromClientTriggerEvents(clientTriggerEvents *client.JobGerritTriggerOnEvents) (*interfaceJobTriggerEvents, error) {
	triggerEvents := interfaceJobTriggerEvents{}
	if clientTriggerEvents != nil && clientTriggerEvents.Items != nil {
		for _, clientTriggerEvent := range *clientTriggerEvents.Items {
			triggerEventType := clientTriggerEvent.GetType()
			initFunc := jobTriggerEventInitFunc[triggerEventType]
			if initFunc == nil {
				return nil, errors.New("Failed to deserialize job gerrit trigger event, missing init func")
			}
			triggerEvent, err := initFunc().fromClientTriggerEvent(clientTriggerEvent)
			if err != nil {
				return nil, err
			}
			mapData := map[string]interface{}{}
			if err := mapstructure.Decode(triggerEvent, &mapData); err != nil {
				return nil, err
			}
			triggerEvents = append(triggerEvents, mapData)
		}
	}
	return &triggerEvents, nil
}
