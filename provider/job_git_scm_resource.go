package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

// ErrJobMultipleDefinitions
var ErrJobMultipleDefinitions = errors.New("Job cannot have have more than one definition")

func jobGitScmResource() *schema.Resource {
	newJobGitScmInterface := func() jobDefinition {
		return newJobGitScm()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobDefinitionCreate(d, m, newJobGitScmInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobDefinitionRead(d, m, newJobGitScmInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobDefinitionUpdate(d, m, newJobGitScmInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobDefinitionDelete(d, m, newJobGitScmInterface)
		},

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
			"config_version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Version",
				Optional:    true,
				Default:     "2",
			},
			"script_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Path to the script to use",
				Optional:    true,
			},
			"lightweight": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If a lightweigth checkout should be done",
				Optional:    true,
				Default:     "false",
			},
		},
	}
}
