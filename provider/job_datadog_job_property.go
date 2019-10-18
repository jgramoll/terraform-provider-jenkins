package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.DatadogJobPropertyType] = func() jobProperty {
		return newJobDatadogJobProperty()
	}
}

type jobDatadogJobProperty struct {
	Type           client.JobPropertyType `mapstructure:"type"`
	Plugin         string                 `mapstructure:"plugin"`
	EmitOnCheckout bool                   `mapstructure:"emit_on_checkout"`
}

func newJobDatadogJobProperty() *jobDatadogJobProperty {
	return &jobDatadogJobProperty{
		Type: client.DatadogJobPropertyType,
	}
}

func (p *jobDatadogJobProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobDatadogJobProperty()
	clientProperty.Plugin = p.Plugin
	clientProperty.EmitOnCheckout = p.EmitOnCheckout
	return clientProperty, nil
}

func (*jobDatadogJobProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobDatadogJobProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobDatadogJobProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobDatadogJobProperty()
	property.Plugin = clientProperty.Plugin
	property.EmitOnCheckout = clientProperty.EmitOnCheckout
	return property, nil
}
