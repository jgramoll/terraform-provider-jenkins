package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyInitFunc[client.DisableConcurrentBuildsJobPropertyType] = func() jobProperty {
		return newJobDisableConcurrentBuildsJobProperty()
	}
}

type jobDisableConcurrentBuildsJobProperty struct {
	Type client.JobPropertyType `mapstructure:"type"`
}

func newJobDisableConcurrentBuildsJobProperty() *jobDisableConcurrentBuildsJobProperty {
	return &jobDisableConcurrentBuildsJobProperty{
		Type: client.DisableConcurrentBuildsJobPropertyType,
	}
}

func (p *jobDisableConcurrentBuildsJobProperty) toClientProperty() (client.JobProperty, error) {
	clientProperty := client.NewJobDisableConcurrentBuildsJobProperty()
	return clientProperty, nil
}

func (p *jobDisableConcurrentBuildsJobProperty) fromClientProperty(clientPropertyInterface client.JobProperty) (jobProperty, error) {
	_, ok := clientPropertyInterface.(*client.JobDisableConcurrentBuildsJobProperty)
	if !ok {
		return nil, fmt.Errorf("Property is not of expected type, expected *client.JobDisableConcurrentBuildsJobProperty, actually %s",
			reflect.TypeOf(clientPropertyInterface).String())
	}
	property := newJobDisableConcurrentBuildsJobProperty()
	return property, nil
}
