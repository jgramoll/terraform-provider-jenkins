package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func jobPropertyResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the property",
				Required:    true,
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[PipelineTriggersJobProperty] Pipeline Trigger",
				Optional:    true,
				Elem:        jobPipelineTriggersPropertyResource(),
			},
			"strategy": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[BuildDiscarderProperty]",
				MaxItems:    1,
				Optional:    true,
				Elem:        jobBuildDiscarderPropertyStrategyResource(),
			},
			"emit_on_checkout": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "[DatadogJobProperty]",
				Optional:    true,
			},
			"parameter": &schema.Schema{
				Type:        schema.TypeList,
				Description: "[ParametersDefinitionProperty]",
				Optional:    true,
				Elem:        jobParametersDefinitionParameterResource(),
			},
		},
	}
}
