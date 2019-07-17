package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// var jobBuildDiscarderPropertyResourceName = "jenkins_job_build_discarder_property_resource"

func jobBuildDiscarderPropertyResource() *schema.Resource {
	newJobBuildDiscarderPropertyInterface := func() jobProperty {
		return newJobBuildDiscarderProperty()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyCreate(d, m, newJobBuildDiscarderPropertyInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyRead(d, m, newJobBuildDiscarderPropertyInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobPropertyDelete(d, m, newJobBuildDiscarderPropertyInterface)
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
