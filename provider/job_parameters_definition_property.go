package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.ParametersDefinitionPropertyType] = func() jobProperty {
		return newJobParametersDefinitionProperty()
	}
}

type jobParametersDefinitionProperty struct {
	Type       client.JobPropertyType  `mapstructure:"type"`
	Parameters *interfaceJobParameters `mapstructure:"parameter"`
}

func newJobParametersDefinitionProperty() *jobParametersDefinitionProperty {
	return &jobParametersDefinitionProperty{
		Type: client.ParametersDefinitionPropertyType,
	}
}

func (p *jobParametersDefinitionProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobParametersDefinitionProperty()

	parameters, err := p.Parameters.toClientParameters()
	if err != nil {
		return nil, err
	}
	clientProperty.ParameterDefinitions = parameters

	return clientProperty, nil
}

func (*jobParametersDefinitionProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobParametersDefinitionProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobParametersDefinitionProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobParametersDefinitionProperty()

	parameters, err := property.Parameters.fromClientParameters(clientProperty.ParameterDefinitions)
	if err != nil {
		return nil, err
	}
	property.Parameters = parameters

	return property, nil
}
