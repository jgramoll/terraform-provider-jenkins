package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type interfaceJobActions []map[string]interface{}

type jobActionInit func() jobAction

var jobActionInitFunc = map[client.JobActionType]jobActionInit{}

type jobAction interface {
	fromClientAction(client.JobAction) (jobAction, error)
	toClientAction(id string) (client.JobAction, error)
	setResourceData(*schema.ResourceData) error
}

func (actions *interfaceJobActions) toClientActions() (*client.JobActions, error) {
	clientActions := client.NewJobActions()
	for _, mapData := range *actions {
		actionTypeString, ok := mapData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize client action, missing type")
		}
		actionType, err := client.ParseJobActionType(actionTypeString)
		if err != nil {
			return nil, err
		}
		action := jobActionInitFunc[actionType]()
		if err := mapstructure.Decode(mapData, &action); err != nil {
			return nil, err
		}
		clientAction, err := action.toClientAction("id")
		if err != nil {
			return nil, err
		}
		*clientActions.Items = append(*clientActions.Items, clientAction)
	}
	return clientActions, nil
}

func (*interfaceJobActions) fromClientActions(clientActions *client.JobActions) (*interfaceJobActions, error) {
	actions := interfaceJobActions{}
	if clientActions != nil && clientActions.Items != nil {
		for _, a := range *clientActions.Items {
			actionType := a.GetType()
			action := jobActionInitFunc[actionType]()
			action, err := action.fromClientAction(a)
			if err != nil {
				return nil, err
			}
			mapData := map[string]interface{}{}
			if err := mapstructure.Decode(action, &mapData); err != nil {
				return nil, err
			}
			mapData["type"] = actionType
			actions = append(actions, mapData)
		}
	}
	return &actions, nil
}
