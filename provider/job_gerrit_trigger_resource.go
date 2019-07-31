package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobGerritTriggerResource() *schema.Resource {
	newJobTriggerInterface := func() jobTrigger {
		return newJobGerritTrigger()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerCreate(d, m, newJobTriggerInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerRead(d, m, newJobTriggerInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerUpdate(d, m, newJobTriggerInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourceJobTriggerDelete(d, m, newJobTriggerInterface)
		},

		Schema: map[string]*schema.Schema{
			"property": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the job trigger property",
				Required:    true,
				ForceNew:    true,
			},
			"server_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the job trigger property",
				Optional:    true,
				Default:     "__ANY__",
			},
			"silent_mode": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Sets silent mode to on or off. When silent mode is on there will be no communication back to Gerrit, i.e. no build started/failed/successful approve messages etc. If other non-silent jobs are triggered by the same Gerrit event as this job, the result of this job's build will not be counted in the end result of the other jobs.",
				Optional:    true,
				Default:     false,
			},
			"silent_start_mode": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "?",
				Optional:    true,
				Default:     false,
			},
			"escape_quotes": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "?",
				Optional:    true,
				Default:     true,
			},
			"name_and_email_parameter_mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "?",
				Optional:    true,
				Default:     "PLAIN",
			},
			"commit_message_parameter_mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the job trigger property",
				Optional:    true,
				Default:     "BASE64",
			},
			"change_subject_parameter_mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "?",
				Optional:    true,
				Default:     "PLAIN",
			},
			"comment_text_parameter_mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "?",
				Optional:    true,
				Default:     "BASE64",
			},
			"dynamic_trigger_configuration": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "?",
				Optional:    true,
				Default:     false,
			},
			"skip_vote": {
				Type:        schema.TypeList,
				Description: "Skip vote",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"on_successful": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"on_failed": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"on_unstable": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"on_not_built": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}
