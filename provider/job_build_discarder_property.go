package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobBuildDiscarderProperty struct{}

func newJobBuildDiscarderProperty() *jobBuildDiscarderProperty {
	return &jobBuildDiscarderProperty{}
}

func (branch *jobBuildDiscarderProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobBuildDiscarderProperty()
	clientProperty.Id = id
	return clientProperty
}

func (branch *jobBuildDiscarderProperty) fromClientJobProperty(clientProperty client.JobProperty) jobProperty {
	property := newJobBuildDiscarderProperty()
	return property
}

func (config *jobBuildDiscarderProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
