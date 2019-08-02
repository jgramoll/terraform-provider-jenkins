package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDeclarativeJobAction struct {
	Plugin string `mapstructure:"plugin"`
}

func newJobDeclarativeJobAction() *jobDeclarativeJobAction {
	return &jobDeclarativeJobAction{}
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

func (a *jobDeclarativeJobAction) toClientAction(id string) (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobAction()
	clientAction.Id = id
	clientAction.Plugin = a.Plugin
	return clientAction, nil
}

func (a *jobDeclarativeJobAction) setResourceData(d *schema.ResourceData) error {
	return d.Set("plugin", a.Plugin)
}
