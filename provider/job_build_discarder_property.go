package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.BuildDiscarderPropertyType] = func() jobProperty {
		return newJobBuildDiscarderProperty()
	}
}

type jobBuildDiscarderProperty struct {
	Type     client.JobPropertyType               `mapstructure:"type"`
	Strategy *jobBuildDiscarderPropertyStrategies `mapstructure:"strategy"`
}

func newJobBuildDiscarderProperty() *jobBuildDiscarderProperty {
	return &jobBuildDiscarderProperty{
		Type: client.BuildDiscarderPropertyType,
	}
}

func (p *jobBuildDiscarderProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobBuildDiscarderProperty()

	strategy, err := p.Strategy.toClientStrategies()
	if err != nil {
		return nil, err
	}
	clientProperty.Strategy = strategy

	return clientProperty, nil
}

func (p *jobBuildDiscarderProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobBuildDiscarderProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobBuildDiscarderProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobBuildDiscarderProperty()

	strategy, err := property.Strategy.fromClientStrategy(clientProperty.Strategy)
	if err != nil {
		return nil, err
	}
	property.Strategy = strategy

	return property, nil
}
