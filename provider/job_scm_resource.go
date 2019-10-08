package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobScmResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the scm type",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name and version of the plugin",
				Optional:    true,
			},
			"config_version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Version",
				Optional:    true,
				Default:     "2",
			},
			"branch": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[GitSCM] Branch to checkout",
				Optional:    true,
				Elem:        jobScmBranchResource(),
			},
			"extension": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[GitSCM] Git extension",
				Optional:    true,
				Elem:        jobGitScmExtensionResource(),
			},
			"user_remote_config": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[GitSCM] User Remote Config",
				Optional:    true,
				Elem:        jobGitScmUserRemoteConfigResource(),
			},
		},
	}
}
