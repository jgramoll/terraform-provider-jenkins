package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobActionInitFunc[client.DeclarativeJobPropertyTrackerActionType] = func() jobAction {
		return newJobDeclarativeJobPropertyTrackerAction()
	}
}

type jobDeclarativeJobPropertyTrackerAction struct {
	Type   client.JobActionType `mapstructure:"type"`
	Plugin string               `mapstructure:"plugin"`
}

func newJobDeclarativeJobPropertyTrackerAction() *jobDeclarativeJobPropertyTrackerAction {
	return &jobDeclarativeJobPropertyTrackerAction{
		Type: client.DeclarativeJobPropertyTrackerActionType,
	}
}

func (*jobDeclarativeJobPropertyTrackerAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	clientAction, ok := clientActionInterface.(*client.JobDeclarativeJobPropertyTrackerAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobPropertyTrackerAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := newJobDeclarativeJobPropertyTrackerAction()
	action.Plugin = clientAction.Plugin
	return action, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) toClientAction() (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	clientAction.Plugin = a.Plugin
	return clientAction, nil
}
