package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerSkipVoteResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"on_successful": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"on_failed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"on_unstable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"on_not_built": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
