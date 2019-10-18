package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobBuildDiscarderPropertyStrategyResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of Build Discarder Property",
				Required:    true,
				ForceNew:    true,
			},
			"days_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "[LogRotator]",
				Optional:    true,
			},
			"num_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "[LogRotator]",
				Optional:    true,
			},
			"artifact_days_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "[LogRotator]",
				Optional:    true,
			},
			"artifact_num_to_keep": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "[LogRotator]",
				Optional:    true,
			},
		},
	}
}
