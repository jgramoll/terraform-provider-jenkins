package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobParametersDefinitionPropertyResource() *schema.Resource {
	newPropertyInterface := func() jobProperty {
		return newJobParametersDefinitionProperty()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyCreate(d, m, newPropertyInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyRead(d, m, newPropertyInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyDelete(d, m, newPropertyInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourceJobPropertyImporter,
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
