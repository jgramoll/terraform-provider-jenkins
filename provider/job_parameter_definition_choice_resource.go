package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobParameterDefinitionChoiceResource() *schema.Resource {
	newParameterDefinitionInterface := func() jobParameterDefinition {
		return newJobParameterDefinitionChoice()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobParameterDefinitionCreate(d, m, newParameterDefinitionInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobParameterDefinitionRead(d, m, newParameterDefinitionInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobParameterDefinitionUpdate(d, m, newParameterDefinitionInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobParameterDefinitionDelete(d, m, newParameterDefinitionInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourceJobParameterDefinitionImporter,
		},

		Schema: map[string]*schema.Schema{
			"property": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the property",
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
			"choices": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of choices",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
