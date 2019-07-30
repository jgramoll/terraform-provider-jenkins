package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerDraftPublishedEventResource() *schema.Resource {
	newTriggerEventInterface := func() jobGerritTriggerEvent {
		return newJobGerritTriggerDraftPublishedEvent()
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

		Schema: map[string]*schema.Schema{
			"trigger": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the trigger",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}
