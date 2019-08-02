package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobDeclarativeJobActionResource() *schema.Resource {
	newJobActionInterface := func() jobAction {
		return newJobDeclarativeJobAction()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobActionCreate(d, m, newJobActionInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobActionRead(d, m, newJobActionInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobActionUpdate(d, m, newJobActionInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobActionDelete(d, m, newJobActionInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourceJobActionImporter,
		},

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Plugin name and version",
				Optional:    true,
			},
		},
	}
}
