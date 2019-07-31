package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"reflect"
)

type jobDeclarativeJobAction struct{}

func newJobDeclarativeJobAction() *jobDeclarativeJobAction {
	return &jobDeclarativeJobAction{}
}

func (a *jobDeclarativeJobAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	_, ok := clientActionInterface.(*client.JobDeclarativeJobAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := newJobDeclarativeJobAction()
	return action, nil
}

func (a *jobDeclarativeJobAction) toClientAction(id string) (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobAction()
	clientAction.Id = id
	return clientAction, nil
}

func (a *jobDeclarativeJobAction) setResourceData(d *schema.ResourceData) error {
	return nil
}
