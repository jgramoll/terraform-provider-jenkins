package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"reflect"
)

type jobBuildDiscarderProperty struct{}

func newJobBuildDiscarderProperty() *jobBuildDiscarderProperty {
	return &jobBuildDiscarderProperty{}
}

func (p *jobBuildDiscarderProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobBuildDiscarderProperty()
	clientProperty.Id = id
	return clientProperty
}

func (p *jobBuildDiscarderProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	_, ok := clientPropertyInterface.(*client.JobBuildDiscarderProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobBuildDiscarderProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobBuildDiscarderProperty()
	return property, nil
}

func (p *jobBuildDiscarderProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
