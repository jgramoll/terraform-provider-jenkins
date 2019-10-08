package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobScmBranchResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the scm branch",
				Required:    true,
			},
		},
	}
}
