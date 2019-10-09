package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.PipelineTriggersJobPropertyType] = func() jobProperty {
		return newJobPipelineTriggersProperty()
	}
}

type jobPipelineTriggersProperty struct {
	Type     client.JobPropertyType `mapstructure:"type"`
	Triggers *jobTriggers           `mapstructure:"trigger"`
}

func newJobPipelineTriggersProperty() *jobPipelineTriggersProperty {
	return &jobPipelineTriggersProperty{
		Type: client.PipelineTriggersJobPropertyType,
	}
}

func (p *jobPipelineTriggersProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobPipelineTriggersProperty()

	triggers, err := p.Triggers.toClientTriggers()
	if err != nil {
		return nil, err
	}
	clientProperty.Triggers = triggers

	return clientProperty, nil
}

func (p *jobPipelineTriggersProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobPipelineTriggersProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobPipelineTriggersProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobPipelineTriggersProperty()

	triggers, err := property.Triggers.fromClientTriggers(clientProperty.Triggers)
	if err != nil {
		return nil, err
	}
	property.Triggers = triggers

	return property, nil
}
