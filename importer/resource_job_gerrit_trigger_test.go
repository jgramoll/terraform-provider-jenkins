package main

import (
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"

	// "github.com/jgramoll/terraform-provider-jenkins/provider"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testJobPipelineTriggersProperties() *[]client.JobProperty {
	gerritBranch := client.NewJobGerritTriggerBranch()
	gerritBranch.CompareType = client.CompareTypeRegExp
	gerritBranch.Pattern = "my-branch"
	gerritFilePath := client.NewJobGerritTriggerFilePath()
	gerritFilePath.CompareType = client.CompareTypeRegExp
	gerritFilePath.Pattern = "my-file-path"
	gerritProject := client.NewJobGerritTriggerProject()
	gerritProject.CompareType = client.CompareTypePlain
	gerritProject.Pattern = "my-project"
	gerritProject.Branches = gerritProject.Branches.Append(gerritBranch)
	gerritProject.FilePaths = gerritProject.FilePaths.Append(gerritFilePath)
	gerritTrigger := client.NewJobGerritTrigger()
	gerritTrigger.Plugin = "gerrit-trigger@2.29.0"
	gerritTrigger.Projects = gerritTrigger.Projects.Append(gerritProject)
	gerritTriggerChangeMergedEvent := client.NewJobGerritTriggerPluginChangeMergedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerChangeMergedEvent)
	gerritTriggerPatchsetEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerPatchsetEvent)
	gerritTriggerDraftEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerDraftEvent)

	triggerJobProperty := client.NewJobPipelineTriggersProperty()
	triggerJobProperty.Triggers = triggerJobProperty.Triggers.Append(gerritTrigger)

	triggerJobProperty2 := client.NewJobPipelineTriggersProperty()
	gerritTrigger2 := client.NewJobGerritTrigger()
	triggerJobProperty2.Triggers = triggerJobProperty2.Triggers.Append(gerritTrigger2)
	return &[]client.JobProperty{triggerJobProperty, triggerJobProperty2}
}

func TestJobGerritTriggerCode(t *testing.T) {
	job := client.NewJob()
	job.Properties.Items = testJobPipelineTriggersProperties()

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
  name     = ""
  plugin   = ""
  disabled = false

  property {
    type = "PipelineTriggersJobProperty"

    trigger {
      type   = "GerritTrigger"
      plugin = "gerrit-trigger@2.29.0"

      server_name       = "__ANY__"
      silent_mode       = false
      silent_start_mode = false
      escape_quotes     = true

      name_and_email_parameter_mode = "PLAIN"
      commit_message_parameter_mode = "BASE64"
      change_subject_parameter_mode = "PLAIN"
      comment_text_parameter_mode   = "BASE64"
      dynamic_trigger_configuration = false

      skip_vote {
        on_successful = false
        on_failed     = false
        on_unstable   = false
        on_not_built  = false
      }

      trigger_on_event {
        type = "PluginChangeMergedEvent"
      }

      trigger_on_event {
        type = "PluginPatchsetCreatedEvent"

        exclude_drafts         = false
        exclude_trivial_rebase = false
        exclude_no_code_change = false
        exclude_private_state  = false
        exclude_wip_state      = false
      }

      trigger_on_event {
        type = "PluginDraftPublishedEvent"
      }

      gerrit_project {
        compare_type = "PLAIN"
        pattern      = "my-project"

        branch {
          compare_type = "REG_EXP"
          pattern      = "my-branch"
        }

        file_path {
          compare_type = "REG_EXP"
          pattern      = "my-file-path"
        }
      }
    }
  }

  property {
    type = "PipelineTriggersJobProperty"

    trigger {
      type   = "GerritTrigger"
      plugin = ""

      server_name       = "__ANY__"
      silent_mode       = false
      silent_start_mode = false
      escape_quotes     = true

      name_and_email_parameter_mode = "PLAIN"
      commit_message_parameter_mode = "BASE64"
      change_subject_parameter_mode = "PLAIN"
      comment_text_parameter_mode   = "BASE64"
      dynamic_trigger_configuration = false

      skip_vote {
        on_successful = false
        on_failed     = false
        on_unstable   = false
        on_not_built  = false
      }
    }
  }
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
