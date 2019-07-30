package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobPipelineTriggersProperty struct{}

func newJobPipelineTriggersProperty() *jobPipelineTriggersProperty {
	return &jobPipelineTriggersProperty{}
}

func (p *jobPipelineTriggersProperty) toClientProperty(id string) client.JobProperty {
	clientProperty := client.NewJobPipelineTriggersProperty()
	clientProperty.Id = id
	return clientProperty
}

func (p *jobPipelineTriggersProperty) fromClientJobProperty(clientProperty client.JobProperty) jobProperty {
	property := newJobPipelineTriggersProperty()
	return property
}

func (p *jobPipelineTriggersProperty) setResourceData(d *schema.ResourceData) error {
	return nil
}
