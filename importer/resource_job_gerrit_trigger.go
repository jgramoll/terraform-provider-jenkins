package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerCodeFuncs["*client.JobGerritTrigger"] = jobGerritTriggerCode
}

func jobGerritTriggerCode(propertyIndex string, triggerIndex string, triggerInterface client.JobTrigger) string {
	trigger := triggerInterface.(*client.JobGerritTrigger)

	triggerOnEvents := jobGerritTriggerOnEventsCode(triggerIndex, trigger.TriggerOnEvents)
	triggerProjects := jobGerritTriggerProjectsCode(propertyIndex, triggerIndex, trigger.Projects)
	dynamicGerritProjects := jobDynamicGerritProjectsCode(trigger.DynamicGerritProjects)
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger" "trigger_%v" {
	property = "${jenkins_job_pipeline_triggers_property.property_%v.id}"

	plugin            = "%v"
	server_name       = "%v"
	silent_mode       = %v
	silent_start_mode = %v
	escape_quotes     = %v

	name_and_email_parameter_mode = "%v"
	commit_message_parameter_mode = "%v"
	change_subject_parameter_mode = "%v"
	comment_text_parameter_mode   = "%v"
	dynamic_trigger_configuration = %v

	skip_vote = {
		on_successful = %v
		on_failed     = %v
		on_unstable   = %v
		on_not_built  = %v
	}
}
`, triggerIndex, propertyIndex,
		trigger.Plugin, trigger.ServerName, trigger.SilentMode, trigger.SilentStartMode, trigger.EscapeQuotes,
		trigger.NameAndEmailParameterMode, trigger.CommitMessageParameterMode, trigger.ChangeSubjectParameterMode,
		trigger.CommentTextParameterMode, trigger.DynamicTriggerConfiguration,
		trigger.SkipVote.OnSuccessful, trigger.SkipVote.OnFailed,
		trigger.SkipVote.OnUnstable, trigger.SkipVote.OnNotBuilt) +
		triggerOnEvents + triggerProjects + dynamicGerritProjects
}
