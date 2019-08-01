package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"reflect"
)

type jobDatadogJobProperty struct {
	EmitOnCheckout bool `mapstruture:"emit_on_checkout"`
}

func newJobDatadogJobProperty() *jobDatadogJobProperty {
	return &jobDatadogJobProperty{}
}

func (p *jobDatadogJobProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobDatadogJobProperty()
	clientProperty.Id = id
	clientProperty.EmitOnCheckout = p.EmitOnCheckout
	return clientProperty
}

func (p *jobDatadogJobProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobDatadogJobProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobDatadogJobProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobDatadogJobProperty()
	property.EmitOnCheckout = clientProperty.EmitOnCheckout
	return property, nil
}

func (p *jobDatadogJobProperty) setResourceData(d *schema.ResourceData) error {
	return d.Set("emit_on_checkout", p.EmitOnCheckout)
}
