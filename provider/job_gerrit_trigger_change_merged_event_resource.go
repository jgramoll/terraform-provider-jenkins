package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerChangeMergedEventResource() *schema.Resource {
	newTriggerEventInterface := func() jobGerritTriggerEvent {
		return newJobGerritTriggerChangeMergedEvent()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventCreate(d, m, newTriggerEventInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventRead(d, m, newTriggerEventInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventDelete(d, m, newTriggerEventInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourceJobTriggerEventImporter,
		},

		Schema: map[string]*schema.Schema{
			"trigger": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the gerrit trigger",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}
