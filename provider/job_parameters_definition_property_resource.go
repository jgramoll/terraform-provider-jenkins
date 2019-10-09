package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobParametersDefinitionParameterResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of trigger",
				Required:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the parameter",
				Required:    true,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Description: "description of the parameter",
				Optional:    true,
			},
			"default_value": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "default value of the parameter",
				Optional:    true,
			},
			"trim": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "trim value of the parameter",
				Optional:    true,
			},
			"choices": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[ChoiceParameterDefinition] List of choices",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
