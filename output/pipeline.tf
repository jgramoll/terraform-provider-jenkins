resource "jenkins_job" "main" {
  name     = "Bridge Career/Platform/Platform_runtime_upload"
  plugin   = "workflow-job@2.33"
  disabled = false

  resource "jenkins_job_pipeline_triggers_property" "property_1" {
    job = "${jenkins_job.main.name}"
  }

  resource "jenkins_job_gerrit_trigger" "trigger_1_1" {
    property = "${jenkins_job_pipeline_triggers_property.property_1.id}"

    plugin            = "gerrit-trigger@2.30.0"
    server_name       = "gerrit.instructure.com"
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

  resource "jenkins_job_gerrit_project" "project_1_1_1" {
    trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

    compare_type = "PLAIN"
    pattern      = "get_smart"
  }

  resource "jenkins_job_gerrit_branch" "branch_1_1_1_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_1.id}"

    compare_type = "REG_EXP"
    pattern      = "^(?!refs/meta/config).*$"
  }

  resource "jenkins_job_gerrit_file_path" "file_path_1_1_1_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_1.id}"

    compare_type = "ANT"
    pattern      = "bridge-platform-runtime.yml"
  }

  resource "jenkins_job_gerrit_project" "project_1_1_2" {
    trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

    compare_type = "PLAIN"
    pattern      = "bridge-platform-runtime"
  }

  resource "jenkins_job_gerrit_branch" "branch_1_1_2_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_2.id}"

    compare_type = "REG_EXP"
    pattern      = "^(?!refs/meta/config).*$"
  }

  resource "jenkins_job_gerrit_file_path" "file_path_1_1_2_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_2.id}"

    compare_type = "ANT"
    pattern      = "bridge-platform-runtime/**"
  }

  resource "jenkins_job_gerrit_file_path" "file_path_1_1_2_2" {
    project = "${jenkins_job_gerrit_project.project_1_1_2.id}"

    compare_type = "ANT"
    pattern      = "bridge-nav/bridge-platform-runtime.yml"
  }

  resource "jenkins_job_gerrit_file_path" "file_path_1_1_2_3" {
    project = "${jenkins_job_gerrit_project.project_1_1_2.id}"

    compare_type = "ANT"
    pattern      = "bridge-status/assets/bridge-platform-runtime.yml"
  }

  resource "jenkins_job_gerrit_project" "project_1_1_3" {
    trigger = "${jenkins_job_gerrit_trigger.trigger_1_1.id}"

    compare_type = "PLAIN"
    pattern      = "bridge-talent-management"
  }

  resource "jenkins_job_gerrit_branch" "branch_1_1_3_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_3.id}"

    compare_type = "REG_EXP"
    pattern      = "^(?!refs/meta/config).*$"
  }

  resource "jenkins_job_gerrit_file_path" "file_path_1_1_3_1" {
    project = "${jenkins_job_gerrit_project.project_1_1_3.id}"

    compare_type = "ANT"
    pattern      = "bridge-talent-web/ui/static/bridge-platform-runtime.yml"
  }

  property {
    type   = "JiraProjectProperty"
    plugin = "jira@3.0.8"
  }

  property {
    type   = "DatadogJobProperty"
    plugin = "datadog@0.7.1"

    emit_on_checkout = false
  }

  property {
    type = "BuildDiscarderProperty"

    strategy {
      type = "LogRotator"

      days_to_keep          = "30"
      num_to_keep           = "100"
      artifact_days_to_keep = "-1"
      artifact_num_to_keep  = "-1"
    }
  }
}

resource "jenkins_job_declarative_job_action" "main" {
  job    = "${jenkins_job.main.name}"
  plugin = "pipeline-model-definition@1.3.9"
}

resource "jenkins_job_declarative_job_property_tracker_action" "main" {
  job    = "${jenkins_job.main.name}"
  plugin = "pipeline-model-definition@1.3.9"
}

resource "jenkins_job_git_scm" "main" {
  job = "${jenkins_job.main.name}"

  plugin     = "workflow-cps@2.70"
  git_plugin = "git@3.10.0"

  config_version = "2"
  script_path    = "bridge-platform-runtime/Jenkinsfile.upload"
  lightweight    = false
}

resource "jenkins_job_git_scm_user_remote_config" "config_1" {
  scm = "${jenkins_job_git_scm.main.id}"

  refspec        = ""
  url            = "ssh://gerrit.instructure.com:29418/bridge-platform-runtime.git"
  credentials_id = "44aa91d6-ab24-498a-b2b4-911bcb17cc35"
}

resource "jenkins_job_git_scm_branch" "branch_1" {
  scm = "${jenkins_job_git_scm.main.id}"

  name = "origin/master"
}
