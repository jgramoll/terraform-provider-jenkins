package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobPipelineTriggersPropertyResource() *schema.Resource {
	newJobGerritPropertyInterface := func() jobProperty {
		return newJobPipelineTriggersProperty()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyCreate(d, m, newJobGerritPropertyInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyRead(d, m, newJobGerritPropertyInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyDelete(d, m, newJobGerritPropertyInterface)
		},

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}