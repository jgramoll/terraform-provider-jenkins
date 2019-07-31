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
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobActionDelete(d, m, newJobActionInterface)
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
