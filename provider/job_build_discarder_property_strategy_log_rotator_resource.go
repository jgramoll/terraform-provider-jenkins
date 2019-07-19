package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobBuildDiscarderPropertyStrategyLogRotatorResource() *schema.Resource {
	newJobBuildDiscarderPropertyStrategyLogRotatorInterface := func() jobBuildDiscarderPropertyStrategy {
		return newJobBuildDiscarderPropertyStrategyLogRotator()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobBuildDiscarderPropertyStrategyCreate(d, m, newJobBuildDiscarderPropertyStrategyLogRotatorInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobBuildDiscarderPropertyStrategyRead(d, m, newJobBuildDiscarderPropertyStrategyLogRotatorInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobBuildDiscarderPropertyStrategyUpdate(d, m, newJobBuildDiscarderPropertyStrategyLogRotatorInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobBuildDiscarderPropertyStrategyDelete(d, m, newJobBuildDiscarderPropertyStrategyLogRotatorInterface)
		},

		Schema: map[string]*schema.Schema{
			"property": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the property",
				Required:    true,
				ForceNew:    true,
			},
			"days_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Number of days to keep job logs",
				Optional:    true,
				Default:     -1,
			},
			"num_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Number of job log to keep",
				Optional:    true,
				Default:     -1,
			},
			"artifact_days_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Number of days to keep job artifacts",
				Optional:    true,
				Default:     -1,
			},
			"artifact_num_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Number of job artifacts to keep",
				Optional:    true,
				Default:     -1,
			},
		},
	}
}
