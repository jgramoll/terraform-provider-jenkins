package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritFilePathResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compare_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of strategy to use for matching gerrit branch",
				Required:    true,
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Pattern to use for matching gerrit branch",
				Required:    true,
			},
		},
	}
}
