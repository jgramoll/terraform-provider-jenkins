package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGitScmUserRemoteConfigResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"refspec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"credentials_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
