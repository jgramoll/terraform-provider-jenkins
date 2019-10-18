package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerEventResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"exclude_drafts": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[PluginPatchsetCreatedEvent]",
				Optional:    true,
			},
			"exclude_trivial_rebase": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[PluginPatchsetCreatedEvent]",
				Optional:    true,
			},
			"exclude_no_code_change": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[PluginPatchsetCreatedEvent]",
				Optional:    true,
			},
			"exclude_private_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[PluginPatchsetCreatedEvent]",
				Optional:    true,
			},
			"exclude_wip_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[PluginPatchsetCreatedEvent]",
				Optional:    true,
			},
		},
	}
}
