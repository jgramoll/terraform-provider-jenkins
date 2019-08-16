package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// Services used by provider
type Services struct {
	Config     client.Config
	JobService client.JobService
}

// Config for provider
type Config struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Token    string `mapstructure:"token"`
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_ADDRESS", nil),
				Description: "Address of jenkins",
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_USERNAME", nil),
				Description: "Name of the user to authenticate with jenkins",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_TOKEN", nil),
				Description: "Token for the user to authenticate with jenkins",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jenkins_job": jobResource(),

			"jenkins_job_declarative_job_action":                  jobDeclarativeJobActionResource(),
			"jenkins_job_declarative_job_property_tracker_action": jobDeclarativeJobPropertyTrackerActionResource(),

			"jenkins_job_git_scm":                                 jobGitScmResource(),
			"jenkins_job_git_scm_user_remote_config":              jobGitScmUserRemoteConfigResource(),
			"jenkins_job_git_scm_branch":                          jobGitScmBranchResource(),
			"jenkins_job_git_scm_clean_before_checkout_extension": jobGitScmCleanBeforeCheckoutExtensionResource(),

			"jenkins_job_build_discarder_property":                      jobBuildDiscarderPropertyResource(),
			"jenkins_job_build_discarder_property_log_rotator_strategy": jobBuildDiscarderPropertyStrategyLogRotatorResource(),

			"jenkins_job_pipeline_triggers_property":            jobPipelineTriggersPropertyResource(),
			"jenkins_job_gerrit_trigger":                        jobGerritTriggerResource(),
			"jenkins_job_gerrit_trigger_change_merged_event":    jobGerritTriggerChangeMergedEventResource(),
			"jenkins_job_gerrit_trigger_patchset_created_event": jobGerritTriggerPatchsetCreatedEventResource(),
			"jenkins_job_gerrit_trigger_draft_published_event":  jobGerritTriggerDraftPublishedEventResource(),
			"jenkins_job_gerrit_project":                        jobGerritProjectResource(),
			"jenkins_job_gerrit_branch":                         jobGerritBranchResource(),
			"jenkins_job_gerrit_file_path":                      jobGerritFilePathResource(),

			"jenkins_job_datadog_job_property":  jobDatadogJobPropertyResource(),
			"jenkins_job_jira_project_property": jobJiraProjectPropertyResource(),

			"jenkins_job_parameters_definition_property": jobParametersDefinitionPropertyResource(),
			"jenkins_job_parameter_definition_choice":    jobParameterDefinitionChoiceResource(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing jenkins client")

	clientConfig := client.Config(config)
	c := client.NewClient(clientConfig)
	return &Services{
		Config:     clientConfig,
		JobService: client.JobService{Client: c},
	}, nil
}
