package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

// type Comparable struct {
// 	CompareType string `mapstructure:"compare_type"`
// 	Pattern     string `mapstructure:"pattern"`
// }

// Job property
type jobPipelineTriggersProperty struct {
	RefId string `mapstructure:"ref_id"`
}

func newJobPipelineTriggersProperty() *jobPipelineTriggersProperty {
	return &jobPipelineTriggersProperty{}
}

func (p *jobPipelineTriggersProperty) setRefID(id string) {
	p.RefId = id
}

func (p *jobPipelineTriggersProperty) getRefID() string {
	return p.RefId
}

func (p *jobPipelineTriggersProperty) toClientProperty() client.JobProperty {
	return &client.JobPipelineTriggersProperty{
		Id: p.RefId,
	}
}

func (p *jobPipelineTriggersProperty) fromClientJobProperty(cs client.JobProperty) jobProperty {
	clientProperty := cs.(*client.JobPipelineTriggersProperty)
	newProperty := newJobPipelineTriggersProperty()

	newProperty.RefId = clientProperty.Id

	return newProperty
}

func (p *jobPipelineTriggersProperty) setResourceData(d *schema.ResourceData) error {
	// err := d.Set("ref_id", p.RefId)
	// if err != nil {
	// 	return err
	// }
	// return d.Set("ref_id", p.RefId)
	return nil
}
