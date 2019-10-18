package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritProjectResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compare_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"branch": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     jobGerritBranchResource(),
			},
			"file_path": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     jobGerritFilePathResource(),
			},
		},
	}
}
