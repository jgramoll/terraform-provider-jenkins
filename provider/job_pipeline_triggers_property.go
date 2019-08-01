package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"reflect"
)

type jobPipelineTriggersProperty struct{}

func newJobPipelineTriggersProperty() *jobPipelineTriggersProperty {
	return &jobPipelineTriggersProperty{}
}

func (p *jobPipelineTriggersProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobPipelineTriggersProperty()
	clientProperty.Id = id
	return clientProperty
}

func (p *jobPipelineTriggersProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	_, ok := clientPropertyInterface.(*client.JobPipelineTriggersProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobPipelineTriggersProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobPipelineTriggersProperty()
	return property, nil
}

func (p *jobPipelineTriggersProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
