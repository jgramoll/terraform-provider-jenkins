package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type jobTriggers []map[string]interface{}

type jobTriggerInit func() jobTrigger

var jobTriggerInitFunc = map[client.JobTriggerType]jobTriggerInit{}

func (triggers *jobTriggers) toClientTriggers() (*client.JobTriggers, error) {
	clientTriggers := client.NewJobTriggers()
	for _, triggerData := range *triggers {
		triggerTypeString, ok := triggerData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize job trigger, missing type")
		}
		triggerType, err := client.ParseJobTriggerType(triggerTypeString)
		if err != nil {
			return nil, err
		}
		initFunc := jobTriggerInitFunc[triggerType]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job trigger, missing init func")
		}
		trigger := initFunc()
		if err := mapstructure.Decode(triggerData, &trigger); err != nil {
			return nil, err
		}
		clientTrigger, err := trigger.toClientTrigger()
		if err != nil {
			return nil, err
		}
		clientTriggers = clientTriggers.Append(clientTrigger)
	}
	return clientTriggers, nil
}

func (*jobTriggers) fromClientTriggers(clientTriggers *client.JobTriggers) (*jobTriggers, error) {
	if clientTriggers == nil || clientTriggers.Items == nil {
		return nil, nil
	}
	triggers := jobTriggers{}

	for _, clientTrigger := range *clientTriggers.Items {
		initFunc := jobTriggerInitFunc[clientTrigger.GetType()]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job trigger, missing init func")
		}
		trigger, err := initFunc().fromClientTrigger(clientTrigger)
		if err != nil {
			return nil, err
		}
		triggerData := map[string]interface{}{}
		if err := mapstructure.Decode(trigger, &triggerData); err != nil {
			return nil, err
		}
		triggers = append(triggers, triggerData)
	}
	return &triggers, nil
}
