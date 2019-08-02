package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDeclarativeJobPropertyTrackerAction struct {
	Plugin string `mapstructure:"plugin"`
}

func newJobDeclarativeJobPropertyTrackerAction() *jobDeclarativeJobPropertyTrackerAction {
	return &jobDeclarativeJobPropertyTrackerAction{}
}

func (a *jobDeclarativeJobPropertyTrackerAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	clientAction, ok := clientActionInterface.(*client.JobDeclarativeJobPropertyTrackerAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobPropertyTrackerAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := newJobDeclarativeJobPropertyTrackerAction()
	action.Plugin = clientAction.Plugin
	return action, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) toClientAction(id string) (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	clientAction.Id = id
	clientAction.Plugin = a.Plugin
	return clientAction, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) setResourceData(d *schema.ResourceData) error {
	return d.Set("plugin", a.Plugin)
}
