package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobActionResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the jenkins action",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name and version of the plugin",
				Optional:    true,
			},
		},
	}
}
