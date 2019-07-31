package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDeclarativeJobPropertyTrackerAction struct{}

func newJobDeclarativeJobPropertyTrackerAction() *jobDeclarativeJobPropertyTrackerAction {
	return &jobDeclarativeJobPropertyTrackerAction{}
}

func (a *jobDeclarativeJobPropertyTrackerAction) fromClientAction(clientActionInterface client.JobAction) (jobAction, error) {
	_, ok := clientActionInterface.(*client.JobDeclarativeJobPropertyTrackerAction)
	if !ok {
		return nil, fmt.Errorf("Failed to parse client action, expected *client.JobDeclarativeJobPropertyTrackerAction, got %s",
			reflect.TypeOf(clientActionInterface).String())
	}
	action := newJobDeclarativeJobPropertyTrackerAction()
	return action, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) toClientAction(id string) (client.JobAction, error) {
	clientAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	clientAction.Id = id
	return clientAction, nil
}

func (a *jobDeclarativeJobPropertyTrackerAction) setResourceData(d *schema.ResourceData) error {
	return nil
}
