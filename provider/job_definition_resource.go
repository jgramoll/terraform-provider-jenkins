package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobDefinitionResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the definition type",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name and id of the plugin",
				Optional:    true,
			},
			"scm": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Job Definition",
				MaxItems:    1,
				Optional:    true,
				Elem:        jobScmResource(),
			},
			"script_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "[CpsScmFlowDefinition] Required. Path to the Script",
				Optional:    true,
			},
			"lightweight": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[CpsScmFlowDefinition] lightweight checkout",
				Optional:    true,
				Default:     false,
			},
		},
	}
}
