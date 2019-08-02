package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGitScmCleanBeforeCheckoutExtensionResource() *schema.Resource {
	newExtensionInterface := func() jobGitScmExtension {
		return newJobGitScmCleanBeforeCheckoutExtension()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobGitScmExtensionCreate(d, m, newExtensionInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobGitScmExtensionRead(d, m, newExtensionInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobGitScmExtensionDelete(d, m, newExtensionInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourceJobGitScmExtensionImporter,
		},

		Schema: map[string]*schema.Schema{
			"scm": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the scm definition",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}
