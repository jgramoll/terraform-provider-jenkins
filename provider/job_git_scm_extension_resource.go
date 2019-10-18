package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGitScmExtensionResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of Extension",
				Required:    true,
			},
		},
	}
}
