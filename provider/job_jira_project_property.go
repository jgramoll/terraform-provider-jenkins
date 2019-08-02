package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobJiraProjectProperty struct {
	Plugin string `mapstructure:"plugin"`
}

func newJobJiraProjectProperty() *jobJiraProjectProperty {
	return &jobJiraProjectProperty{}
}

func (p *jobJiraProjectProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobJiraProjectProperty()
	clientProperty.Id = id
	clientProperty.Plugin = p.Plugin
	return clientProperty
}

func (p *jobJiraProjectProperty) fromClientJobProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobJiraProjectProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobJiraProjectProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobJiraProjectProperty()
	property.Plugin = clientProperty.Plugin
	return property, nil
}

func (p *jobJiraProjectProperty) setResourceData(d *schema.ResourceData) error {
	return d.Set("plugin", p.Plugin)
}
