package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobParametersDefinitionProperty struct {
}

func newJobParametersDefinitionProperty() *jobParametersDefinitionProperty {
	return &jobParametersDefinitionProperty{}
}

func (p *jobParametersDefinitionProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobParametersDefinitionProperty()
	clientProperty.Id = id
	return clientProperty
}

func (p *jobParametersDefinitionProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	_, ok := clientPropertyInterface.(*client.JobParametersDefinitionProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobParametersDefinitionProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobParametersDefinitionProperty()
	return property, nil
}

func (p *jobParametersDefinitionProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
