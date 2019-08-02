package client

import (
	"reflect"
	"strings"
	"testing"
)

var jobService *JobService

func init() {
	jobService = &JobService{newTestClient()}
}

func TestGetJobs(t *testing.T) {
	jobs, err := jobService.GetJobs("Bridge Career")
	if err != nil {
		t.Fatal(err)
	}

	if len(*jobs) == 0 {
		t.Fatal("Should have jobs")
	}
}

func TestGetJobDetails(t *testing.T) {
	jobName := "Bridge Career/migrations_change"
	job, err := jobService.GetJob(jobName)
	if err != nil {
		t.Fatal(err)
	}
	if job.Name != "Bridge Career/migrations_change" {
		t.Fatalf("Job name should be %v, was %v", "Bridge Career/migrations_change", job.Name)
	}
	if job.Description != "" {
		t.Fatalf("Job description should be %v, was %v", "", job.Description)
	}
	properties := *job.Properties.Items
	if len(properties) != 4 {
		t.Fatalf("Job should have %v properties, was %v", 4, len(properties))
	}
	pipelineTriggersProperty, ok := properties[0].(*JobPipelineTriggersProperty)
	if !ok {
		t.Fatalf("Invalid pipeline triggers property, got %s", reflect.TypeOf(properties[0]).String())
	}
	_, ok = properties[1].(*JobJiraProjectProperty)
	if !ok {
		t.Fatalf("Invalid jira project property, got %s", reflect.TypeOf(properties[1]).String())
	}
	_, ok = properties[2].(*JobDatadogJobProperty)
	if !ok {
		t.Fatalf("Invalid datadog property, got %s", reflect.TypeOf(properties[2]).String())
	}
	pipelineDiscarderProperty, ok := properties[3].(*JobBuildDiscarderProperty)
	if !ok {
		t.Fatalf("Invalid build discarder property, got %s", reflect.TypeOf(properties[3]).String())
	}
	if pipelineDiscarderProperty.Strategy.Item == nil {
		t.Fatalf("Job missing discarder strategy")
	}
	triggers := *pipelineTriggersProperty.Triggers
	if len(*triggers.Items) != 1 {
		t.Fatalf("Job should have %v triggers, was %v", 1, len(*triggers.Items))
	}
	gerritTrigger, ok := (*triggers.Items)[0].(*JobGerritTrigger)
	if !ok {
		t.Fatalf("Job should have %v, was %v", "*client.JobGerritTrigger", reflect.TypeOf((*triggers.Items)[0]))
	}
	if gerritTrigger.ServerName != "gerrit.instructure.com" {
		t.Fatalf("Job Trigger ServerName should be %v, was %v", "gerrit.instructure.com", gerritTrigger.ServerName)
	}
	projects := *(gerritTrigger.Projects).Items
	if len(projects) != 1 {
		t.Fatalf("Job should have %v trigger gerrit projects, was %v", 1, len(projects))
	}
	gerritProject := projects[0]
	if gerritProject.CompareType != CompareTypePlain {
		t.Fatalf("Job should have gerrit project compare type %v, was %v", CompareTypePlain, gerritProject.CompareType)
	}
	if gerritProject.Pattern != "bridge-career-infrastructure" {
		t.Fatalf("Job should have gerrit project pattern %v, was %v", "bridge-career-infrastructure", gerritProject.Pattern)
	}
	branches := *(gerritProject.Branches).Items
	if len(branches) != 1 {
		t.Fatalf("Job should have %v trigger gerrit branches, was %v", 1, len(branches))
	}
	branch := branches[0]
	if branch.CompareType != CompareTypeRegExp {
		t.Fatalf("Job should have gerrit branch compare type %v, was %v", CompareTypeRegExp, branch.CompareType)
	}
	if branch.Pattern != "^(?!refs/meta/config).*$" {
		t.Fatalf("Job should have gerrit branch pattern %v, was %v", "^(?!refs/meta/config).*$", branch.Pattern)
	}
	definition := job.Definition.(*CpsScmFlowDefinition)
	if definition.ScriptPath != "migrations.Jenkinsfile" {
		t.Fatalf("Job should have script path %v, was %v", "migrations.Jenkinsfile", definition.ScriptPath)
	}
	scm := *definition.SCM
	userRemoteConfigs := *(scm.UserRemoteConfigs).Items
	if len(userRemoteConfigs) != 1 {
		t.Fatalf("Job should have %v user remote configs, was %v", 1, len(userRemoteConfigs))
	}
	userRemoteConfig := userRemoteConfigs[0]
	if userRemoteConfig.Url != "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git" {
		t.Fatalf("Job should have user remote config url %v, was %v", "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git", userRemoteConfig.Url)
	}
}

func TestCreateJob(t *testing.T) {
	job := NewJob()
	job.Name = "Bridge Career/my test job"
	newItems := append(*(job.Properties).Items, &JobPipelineTriggersProperty{})
	job.Properties.Items = &newItems

	var err error
	err = jobService.CreateJob(job)
	if err != nil {
		t.Fatal(err)
	}

	job.Description = "my new desc 4"
	err = jobService.UpdateJob(job)
	if err != nil {
		t.Fatal(err)
	}

	err = jobService.DeleteJob(job.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCleanup(t *testing.T) {
	jobs, err := jobService.GetJobs("Bridge Career")
	if err != nil {
		t.Fatal(err)
	}

	for _, job := range *jobs {
		if strings.Contains(job.Name, "tf-acc") {
			jobService.DeleteJob(job.Name)
		}
	}
}
