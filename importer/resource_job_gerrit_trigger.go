package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	ensureJobTriggerFuncs["*client.JobGerritTrigger"] = ensureJobGerritTrigger
	jobTriggerCodeFuncs["*client.JobGerritTrigger"] = jobGerritTriggerCode
	jobTriggerImportScriptFuncs["*client.JobGerritTrigger"] = jobGerritTriggerImportScript
}

func ensureJobGerritTrigger(triggerInterface client.JobTrigger) error {
	trigger := triggerInterface.(*client.JobGerritTrigger)
	if trigger.Id == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		trigger.Id = id.String()
	}
	if err := ensureJobDynamicGerritProjects(trigger.DynamicGerritProjects); err != nil {
		return err
	}
	if err := ensureJobGerritTriggerProjects(trigger.Projects); err != nil {
		return err
	}
	return ensureJobGerritTriggerOnEvents(trigger.TriggerOnEvents)
}

func jobGerritTriggerCode(triggerInterface client.JobTrigger) string {
	trigger := triggerInterface.(*client.JobGerritTrigger)

	triggerOnEvents := jobGerritTriggerOnEventsCode(trigger.TriggerOnEvents)
	triggerProjects := jobGerritTriggerProjectsCode(trigger.Projects)
	dynamicGerritProjects := jobDynamicGerritProjectsCode(trigger.DynamicGerritProjects)
	return fmt.Sprintf(`
resource "jenkins_job_gerrit_trigger" "main" {
	property = "${jenkins_job_pipeline_triggers_property.main.id}"

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
`, trigger.Plugin, trigger.ServerName, trigger.SilentMode, trigger.SilentStartMode, trigger.EscapeQuotes,
		trigger.NameAndEmailParameterMode, trigger.CommitMessageParameterMode, trigger.ChangeSubjectParameterMode,
		trigger.CommentTextParameterMode, trigger.DynamicTriggerConfiguration,
		trigger.SkipVote.OnSuccessful, trigger.SkipVote.OnFailed,
		trigger.SkipVote.OnUnstable, trigger.SkipVote.OnNotBuilt) +
		triggerOnEvents + triggerProjects + dynamicGerritProjects
}

func jobGerritTriggerImportScript(jobName string, propertyId string, triggerInterface client.JobTrigger) string {
	trigger := triggerInterface.(*client.JobGerritTrigger)

	return fmt.Sprintf(`
terraform import jenkins_job_gerrit_trigger.main "%v"
`, provider.ResourceJobTriggerId(jobName, propertyId, trigger.Id)) +
		jobGerritTriggerOnEventsImportScript(jobName, propertyId, trigger.Id, trigger.TriggerOnEvents) +
		jobGerritTriggerProjectsImportScript(jobName, propertyId, trigger.Id, trigger.Projects) +
		jobDynamicGerritProjectsImportScript(trigger.DynamicGerritProjects)
}
