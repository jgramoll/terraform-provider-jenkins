package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.JiraProjectPropertyType] = func() jobProperty {
		return newJobJiraProjectProperty()
	}
}

type jobJiraProjectProperty struct {
	Type   client.JobPropertyType `mapstructure:"type"`
	Plugin string                 `mapstructure:"plugin"`
}

func newJobJiraProjectProperty() *jobJiraProjectProperty {
	return &jobJiraProjectProperty{
		Type: client.JiraProjectPropertyType,
	}
}

func (p *jobJiraProjectProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobJiraProjectProperty()
	clientProperty.Plugin = p.Plugin
	return clientProperty, nil
}

func (p *jobJiraProjectProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	clientProperty, ok := clientPropertyInterface.(*client.JobJiraProjectProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobJiraProjectProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobJiraProjectProperty()
	property.Plugin = clientProperty.Plugin
	return property, nil
}
