package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobPipelineJobGerritTriggerResource() *schema.Resource {
	newJobGerritTriggerInterface := func() jobTrigger {
		return newJobPipelineTriggersProperty()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerCreate(d, m, newJobGerritTriggerInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerRead(d, m, newJobGerritTriggerInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerUpdate(d, m, newJobGerritTriggerInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerDelete(d, m, newJobGerritTriggerInterface)
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
