package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerPatchSetCreatedEventResource() *schema.Resource {
	newTriggerEventInterface := func() jobGerritTriggerEvent {
		return newJobGerritTriggerPatchSetCreatedEvent()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventCreate(d, m, newTriggerEventInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventRead(d, m, newTriggerEventInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerEventUpdate(d, m, newTriggerEventInterface)
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
			},
			"exclude_drafts": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If drafts should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_trivial_rebase": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If trivial rebase should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_no_code_change": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If no code change should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_private_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If private state should be considered for triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_wip_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If wip should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
		},
	}
}
