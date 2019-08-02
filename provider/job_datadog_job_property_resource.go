package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobDatadogJobPropertyResource() *schema.Resource {
	newPropertyInterface := func() jobProperty {
		return newJobDatadogJobProperty()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyCreate(d, m, newPropertyInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyRead(d, m, newPropertyInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyUpdate(d, m, newPropertyInterface)
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
			"emit_on_checkout": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Emit on Checkout",
				Optional:    true,
				Default:     false,
			},
		},
	}
}
