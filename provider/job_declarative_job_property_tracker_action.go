package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDeclarativeJobPropertyTrackerAction struct {
	Type   string `mapstructure:"type"`
	Plugin string `mapstructure:"plugin"`
}

func newJobDeclarativeJobPropertyTrackerAction() jobAction {
	return &jobDeclarativeJobPropertyTrackerAction{
		Type: string(client.DeclarativeJobPropertyTrackerActionType),
	}
}

func (a *jobDeclarativeJobPropertyTrackerAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	clientAction, ok := clientActionInterface.(*client.JobDeclarativeJobPropertyTrackerAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobPropertyTrackerAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := jobDeclarativeJobPropertyTrackerAction{}
	action.Plugin = clientAction.Plugin
	return &action, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) toClientAction() (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	clientAction.Plugin = a.Plugin
	return clientAction, nil
}
