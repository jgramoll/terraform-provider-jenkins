package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobTriggerCodeFuncs["*client.JobGerritTrigger"] = jobGerritTriggerCode
}

func jobGerritTriggerCode(triggerInterface client.JobTrigger) string {
	trigger := triggerInterface.(*client.JobGerritTrigger)
	return fmt.Sprintf(`
    trigger {
      type   = "GerritTrigger"
      plugin = "%s"

      server_name       = "%s"
      silent_mode       = %v
      silent_start_mode = %v
      escape_quotes     = %v

      name_and_email_parameter_mode = "%s"
      commit_message_parameter_mode = "%s"
      change_subject_parameter_mode = "%s"
      comment_text_parameter_mode   = "%s"
      dynamic_trigger_configuration = %v

      skip_vote {
        on_successful = %v
        on_failed     = %v
        on_unstable   = %v
        on_not_built  = %v
      }
%s%s    }
`, trigger.Plugin, trigger.ServerName, trigger.SilentMode, trigger.SilentStartMode, trigger.EscapeQuotes,
		trigger.NameAndEmailParameterMode, trigger.CommitMessageParameterMode, trigger.ChangeSubjectParameterMode,
		trigger.CommentTextParameterMode, trigger.DynamicTriggerConfiguration,
		trigger.SkipVote.OnSuccessful, trigger.SkipVote.OnFailed,
		trigger.SkipVote.OnUnstable, trigger.SkipVote.OnNotBuilt,
		jobGerritTriggerOnEventsCode(trigger.TriggerOnEvents),
		jobGerritTriggerProjectsCode(trigger.Projects))
}
