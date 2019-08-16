package main

import (
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testNewJob() *client.Job {
	job := client.NewJob()
	job.Id = "jobId"
	job.Name = "Premerge checks"
	job.Plugin = "flow-plugin"
	job.Description = "my-desc"
	declarativeJobAction := client.NewJobDeclarativeJobAction()
	declarativeJobAction.Id = "declarativeJobActionId"
	declarativeJobAction.Plugin = "DeclarativeJobPlugin"
	job.Actions = job.Actions.Append(declarativeJobAction)
	declarativeJobPropertyTrackerAction := client.NewJobDeclarativeJobPropertyTrackerAction()
	declarativeJobPropertyTrackerAction.Id = "declarativeJobPropertyTrackerActionId"
	declarativeJobPropertyTrackerAction.Plugin = "DeclarativeJobTrackerPlugin"
	job.Actions = job.Actions.Append(declarativeJobPropertyTrackerAction)

	definition := client.NewCpsScmFlowDefinition()
	definition.Id = "definitionId"
	definition.Plugin = "gitPlugin"
	definition.ScriptPath = "my-Jenkinsfile"
	definition.SCM = client.NewGitScm()
	definition.SCM.Plugin = "gitScmPlugin"
	definition.SCM.ConfigVersion = "my-version"

	remoteConfig := client.NewGitUserRemoteConfig()
	remoteConfig.Id = "remoteConfigId"
	remoteConfig.Refspec = "${GERRIT_REFSPEC}"
	remoteConfig.Url = "url.to.server"
	remoteConfig.CredentialsId = "creds"
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(remoteConfig)

	scmExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	scmExtension.Id = "scmExtensionId"
	scmExtension.Id = "extension-id"
	definition.SCM.Extensions = definition.SCM.Extensions.Append(scmExtension)

	branchSpec := client.NewGitScmBranchSpec()
	branchSpec.Id = "branchspecId"
	branchSpec.Name = "branchspec"
	definition.SCM.Branches = definition.SCM.Branches.Append(branchSpec)
	job.Definition = definition

	gerritBranch := client.NewJobGerritTriggerBranch()
	gerritBranch.Id = "gerritBranchId"
	gerritBranch.CompareType = client.CompareTypeRegExp
	gerritBranch.Pattern = "my-branch"
	gerritFilePath := client.NewJobGerritTriggerFilePath()
	gerritFilePath.Id = "gerritFilePathId"
	gerritFilePath.CompareType = client.CompareTypeRegExp
	gerritFilePath.Pattern = "my-file-path"
	gerritProject := client.NewJobGerritTriggerProject()
	gerritProject.Id = "gerritProjectId"
	gerritProject.CompareType = client.CompareTypePlain
	gerritProject.Pattern = "my-project"
	gerritProject.Branches = gerritProject.Branches.Append(gerritBranch)
	gerritProject.FilePaths = gerritProject.FilePaths.Append(gerritFilePath)
	gerritTrigger := client.NewJobGerritTrigger()
	gerritTrigger.Id = "gerrit-trigger-id"
	gerritTrigger.Plugin = "gerrit-trigger@2.29.0"
	gerritTrigger.Projects = gerritTrigger.Projects.Append(gerritProject)
	gerritTriggerChangeMergedEvent := client.NewJobGerritTriggerPluginChangeMergedEvent()
	gerritTriggerChangeMergedEvent.Id = "gerritTriggerMergeEventId"
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerChangeMergedEvent)
	gerritTriggerPatchsetEvent := client.NewJobGerritTriggerPluginPatchsetCreatedEvent()
	gerritTriggerPatchsetEvent.Id = "gerritTriggerPatchsetEventId"
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerPatchsetEvent)
	gerritTriggerDraftEvent := client.NewJobGerritTriggerPluginDraftPublishedEvent()
	gerritTriggerDraftEvent.Id = "gerritTriggerDraftEventId"
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerDraftEvent)
	triggerJobProperty := client.NewJobPipelineTriggersProperty()
	triggerJobProperty.Id = "trigger-id"
	triggerJobProperty.Triggers = triggerJobProperty.Triggers.Append(gerritTrigger)
	job.Properties = job.Properties.Append(triggerJobProperty)

	discardPropertyStrategy := client.NewJobBuildDiscarderPropertyStrategyLogRotator()
	discardPropertyStrategy.Id = "discardPropertyStrategyId"
	discardPropertyStrategy.DaysToKeep = 1
	discardPropertyStrategy.NumToKeep = 2
	discardPropertyStrategy.ArtifactDaysToKeep = 3
	discardPropertyStrategy.ArtifactNumToKeep = 4
	discardProperty := client.NewJobBuildDiscarderProperty()
	discardProperty.Id = "discard-id"
	discardProperty.Strategy.Item = discardPropertyStrategy
	job.Properties = job.Properties.Append(discardProperty)

	datadogJobProperty := client.NewJobDatadogJobProperty()
	datadogJobProperty.Id = "datadogJobPropertyId"
	datadogJobProperty.Plugin = "datadog@0.7.1"
	job.Properties = job.Properties.Append(datadogJobProperty)

	jiraProjectProperty := client.NewJobJiraProjectProperty()
	jiraProjectProperty.Id = "jiraProjectPropertyId"
	jiraProjectProperty.Plugin = "jiraPlugin"
	job.Properties = job.Properties.Append(jiraProjectProperty)
	return job
}

func TestJobCode(t *testing.T) {
	job := testNewJob()

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
	name     = "Premerge checks"
	plugin   = "flow-plugin"
	disabled = false
}

resource "jenkins_job_declarative_job_action" "main" {
	job = "${jenkins_job.main.name}"
	plugin = "DeclarativeJobPlugin"
}

resource "jenkins_job_declarative_job_property_tracker_action" "main" {
	job = "${jenkins_job.main.name}"
	plugin = "DeclarativeJobTrackerPlugin"
}

resource "jenkins_job_git_scm" "main" {
	job = "${jenkins_job.main.name}"

	plugin     = "gitPlugin"
	git_plugin = "gitScmPlugin"

	config_version = "my-version"
	script_path    = "my-Jenkinsfile"
	lightweight    = false
}

resource "jenkins_job_git_scm_user_remote_config" "config_1" {
	scm = "${jenkins_job_git_scm.main.id}"

	refspec        = "$${GERRIT_REFSPEC}"
	url            = "url.to.server"
	credentials_id = "creds"
}

resource "jenkins_job_git_scm_branch" "branch_1" {
	scm = "${jenkins_job_git_scm.main.id}"

	name = "branchspec"
}

resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
	scm = "${jenkins_job_git_scm.main.id}"
}

resource "jenkins_job_pipeline_triggers_property" "property_1" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_gerrit_trigger" "trigger_1_1" {
	property = "${jenkins_job_pipeline_triggers_property.property_1.id}"

	plugin            = "gerrit-trigger@2.29.0"
	server_name       = "__ANY__"
	silent_mode       = false
	silent_start_mode = false
	escape_quotes     = true

	name_and_email_parameter_mode = "PLAIN"
	commit_message_parameter_mode = "BASE64"
	change_subject_parameter_mode = "PLAIN"
	comment_text_parameter_mode   = "BASE64"
	dynamic_trigger_configuration = false

	skip_vote = {
		on_successful = false
		on_failed     = false
		on_unstable   = false
		on_not_built  = false
	}
}

resource "jenkins_job_gerrit_trigger_change_merged_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"
}

resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

	exclude_drafts         = false
	exclude_trivial_rebase = false
	exclude_no_code_change = false
	exclude_private_state  = false
	exclude_wip_state      = false
}

resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"
}

resource "jenkins_job_gerrit_project" "project_1" {
	trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

	compare_type = "PLAIN"
	pattern      = "my-project"
}

resource "jenkins_job_gerrit_branch" "branch_1_1" {
	project = "${jenkins_job_gerrit_project.project_1.id}"

	compare_type = "REG_EXP"
	pattern      = "my-branch"
}

resource "jenkins_job_gerrit_file_path" "file_path_1_1" {
	project = "${jenkins_job_gerrit_project.project_1.id}"

	compare_type = "REG_EXP"
	pattern      = "my-file-path"
}

resource "jenkins_job_build_discarder_property" "property_2" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
	property = "${jenkins_job_build_discarder_property.property_2.id}"

	days_to_keep          = "1"
	num_to_keep           = "2"
	artifact_days_to_keep = "3"
	artifact_num_to_keep  = "4"
}

resource "jenkins_job_datadog_job_property" "property_3" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_jira_project_property" "property_4" {
	job = "${jenkins_job.main.name}"

	plugin = "jiraPlugin"
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestJobImportScript(t *testing.T) {
	job := testNewJob()

	result := jobImportScript(job)
	expected := `terraform init

terraform import jenkins_job.main "Premerge checks"

terraform import jenkins_job_declarative_job_action.main "Premerge checksdeclarativeJobActionId"

terraform import jenkins_job_declarative_job_property_tracker_action.main "Premerge checksdeclarativeJobPropertyTrackerActionId"

terraform import jenkins_job_git_scm.main "Premerge checksdefinitionId"

terraform import jenkins_job_git_scm_user_remote_config.config_1 "Premerge checksdefinitionIdremoteConfigId"

terraform import jenkins_job_git_scm_branch.branch_1 "Premerge checksdefinitionIdbranchspecId"

terraform import jenkins_job_git_scm_clean_before_checkout_extension.main "Premerge checksdefinitionIdextension-id"

terraform import jenkins_job_pipeline_triggers_property.property_1 "Premerge checkstrigger-id"

terraform import jenkins_job_gerrit_trigger.trigger_1_1 "Premerge checkstrigger-idgerrit-trigger-id"

terraform import jenkins_job_gerrit_trigger_change_merged_event.main "Premerge checkstrigger-idgerrit-trigger-idgerritTriggerMergeEventId"

terraform import jenkins_job_gerrit_trigger_patchset_created_event.main "Premerge checkstrigger-idgerrit-trigger-idgerritTriggerPatchsetEventId"

terraform import jenkins_job_gerrit_trigger_draft_published_event.main "Premerge checkstrigger-idgerrit-trigger-idgerritTriggerDraftEventId"

terraform import jenkins_job_gerrit_project.project_1 "Premerge checkstrigger-idgerrit-trigger-idgerritProjectId"

terraform import jenkins_job_gerrit_branch.branch_1_1 "Premerge checkstrigger-idgerrit-trigger-idgerritProjectIdgerritBranchId"

terraform import jenkins_job_gerrit_file_path.file_path_1_1 "Premerge checkstrigger-idgerrit-trigger-idgerritProjectIdgerritFilePathId"

terraform import jenkins_job_build_discarder_property.property_2 "Premerge checksdiscard-id"

terraform import jenkins_job_build_discarder_property_log_rotator_strategy.main "Premerge checksdiscard-iddiscardPropertyStrategyId"

terraform import jenkins_job_datadog_job_property.property_3 "Premerge checksdatadogJobPropertyId"

terraform import jenkins_job_jira_project_property.property_4 "Premerge checksjiraProjectPropertyId"
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform import script not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
