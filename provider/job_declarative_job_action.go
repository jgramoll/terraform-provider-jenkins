package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobActionInitFunc[client.DeclarativeJobActionType] = func() jobAction {
		return newJobDeclarativeJobAction()
	}
}

type jobDeclarativeJobAction struct {
	Type   string `mapstructure:"type"`
	Plugin string `mapstructure:"plugin"`
}

func newJobDeclarativeJobAction() *jobDeclarativeJobAction {
	return &jobDeclarativeJobAction{
		Type: string(client.DeclarativeJobActionType),
	}
}

func (a *jobDeclarativeJobAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	clientAction, ok := clientActionInterface.(*client.JobDeclarativeJobAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := newJobDeclarativeJobAction()
	action.Plugin = clientAction.Plugin
	return action, nil
}

func (a *jobDeclarativeJobAction) toClientAction() (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobAction()
	clientAction.Plugin = a.Plugin
	return clientAction, nil
}
