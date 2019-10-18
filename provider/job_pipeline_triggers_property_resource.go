package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobPipelineTriggersPropertyResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of trigger",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Trigger plugin",
				Optional:    true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "__ANY__",
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
				Elem:        jobGerritTriggerSkipVoteResource(),
			},
			"gerrit_project": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[GerritTrigger]",
				Optional:    true,
				Elem:        jobGerritProjectResource(),
			},
			"trigger_on_event": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[GerritTrigger]",
				Optional:    true,
				Elem:        jobGerritTriggerEventResource(),
			},
		},
	}
}
