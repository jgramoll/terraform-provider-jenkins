# terraform-provider-jenkins
Terraform Provider to manage jenkins jobs

## Install ##

You will need to install the binary as a [terraform third party plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).  Terraform will then pick up the binary from the local filesystem when you run `terraform init`.

```sh
curl -s https://raw.githubusercontent.com/jgramoll/terraform-provider-jenkins/master/install.sh | bash
```

## Usage ##

### Credentials ###

go to $jenkins_url/user/$username/configure

Add to .bash_profile

```sh
export JENKINS_ADDRESS=https://your.jenkins.server
export JENKINS_USERNAME=username
export JENKINS_TOKEN=jenkins_token_from_configure
```

### Importer ###

```sh
go get github.com/jgramoll/terraform-provider-jenkins/importer
importer --job "Name of your Job" --output my_tf_dir
```

1. Ensures each resource has a valid id
1. Outputs terraform code to match the job
1. Outputs script that will import the resources to tf state

### Resources ###

```terraform
provider "jenkins" {
  address = "${var.jenkins_address}"
}

resource "jenkins_job" "main" {
  name     = "Premerge checks"
  plugin   = "workflow-job@2.33"
  disabled = false
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
  script_path    = "Jenkinsfile.api"
  lightweight    = false
}

resource "jenkins_job_git_scm_user_remote_config" "config_1" {
  scm = "${jenkins_job_git_scm.main.id}"

  refspec        = "$${GERRIT_REFSPEC}"
  url            = "ssh://git.server/git-repo.git"
  credentials_id = "123-abc"
}

resource "jenkins_job_git_scm_branch" "branch_1" {
  scm = "${jenkins_job_git_scm.main.id}"

  name = "FETCH_HEAD"
}

resource "jenkins_job_git_scm_clean_before_checkout_extension" "main" {
  scm = "${jenkins_job_git_scm.main.id}"
}

resource "jenkins_job_pipeline_triggers_property" "main" {
  job = "${jenkins_job.main.name}"
}

resource "jenkins_job_gerrit_trigger" "main" {
  property = "${jenkins_job_pipeline_triggers_property.main.id}"

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
  trigger = "${jenkins_job_gerrit_trigger.main.id}"
}

resource "jenkins_job_gerrit_trigger_patchset_created_event" "main" {
  trigger = "${jenkins_job_gerrit_trigger.main.id}"

  exclude_drafts         = false
  exclude_trivial_rebase = false
  exclude_no_code_change = false
  exclude_private_state  = false
  exclude_wip_state      = false
}

resource "jenkins_job_gerrit_trigger_draft_published_event" "main" {
  trigger = "${jenkins_job_gerrit_trigger.main.id}"
}

resource "jenkins_job_gerrit_project" "project_1" {
  trigger = "${jenkins_job_gerrit_trigger.main.id}"

  compare_type = "PLAIN"
  pattern      = "bridge-skills"
}

resource "jenkins_job_gerrit_branch" "branch_1" {
  project = "${jenkins_job_gerrit_project.project_1.id}"

  compare_type = "REG_EXP"
  pattern      = "^(?!refs/meta/config).*$"
}

resource "jenkins_job_gerrit_file_path" "file_path_1" {
  project = "${jenkins_job_gerrit_project.project_1.id}"

  compare_type = "ANT"
  pattern      = "path/to/file"
}

resource "jenkins_job_datadog_job_property" "main" {
  job = "${jenkins_job.main.name}"
}

resource "jenkins_job_build_discarder_property" "main" {
  job = "${jenkins_job.main.name}"
}

resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
  property = "${jenkins_job_build_discarder_property.main.id}"

  days_to_keep          = "30"
  num_to_keep           = "-1"
  artifact_days_to_keep = "-1"
  artifact_num_to_keep  = "-1"
}

resource "jenkins_job_datadog_job_property" "main" {
  job = "${jenkins_job.main.name}"
}

resource "jenkins_job_jira_project_property" "main" {
  job = "${jenkins_job.main.name}"

  plugin = "jiraPlugin"
}

resource "jenkins_job_parameters_definition_property" "parameters" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_parameter_definition_choice" "env" {
	property = "${jenkins_job_parameters_definition_property.parameters.id}"

	name = "env"
	description = "which env to target"
	choices = ["1", "3", "4"]
}

```

## Development ##

### Dependencies ###

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

[Go modules](https://github.com/golang/go/wiki/Modules) are used for dependency management.  To install all dependencies run the following:

```sh
export GO111MODULE=on
go mod vendor
```

### Link ###

```sh
go clean
go build
rm ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-jenkins_v1.0.0
ln  ./terraform-provider-jenkins ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-jenkins_v1.0.0
```
