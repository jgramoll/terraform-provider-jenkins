package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobDisableConcurrentBuildsJobProperty struct{}

func newJobDisableConcurrentBuildsJobProperty() *jobDisableConcurrentBuildsJobProperty {
	return &jobDisableConcurrentBuildsJobProperty{}
}

func (p *jobDisableConcurrentBuildsJobProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobDisableConcurrentBuildsJobProperty()
	clientProperty.Id = id
	return clientProperty
}

func (p *jobDisableConcurrentBuildsJobProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	_, ok := clientPropertyInterface.(*client.JobDisableConcurrentBuildsJobProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobDisableConcurrentBuildsJobProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobDisableConcurrentBuildsJobProperty()
	return property, nil
}

func (p *jobDisableConcurrentBuildsJobProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
