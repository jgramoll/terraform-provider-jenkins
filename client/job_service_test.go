package client

import (
	"testing"
)

var jobService *JobService

const folder string = "job/Bridge%20Career"

func init() {
	// rand.Seed(time.Now().UTC().UnixNano())
	jobService = &JobService{newTestClient()}
}

func TestGetJobs(t *testing.T) {
	jobs, err := jobService.GetJobs()
	if err != nil {
		t.Fatal(err)
	}

	if len(*jobs) == 0 {
		t.Fatal("Should have jobs")
	}
}

func TestGetJob(t *testing.T) {
	jobConfig, err := jobService.GetJob(folder, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "Bridge Career/migrations"
	if jobConfig.FullName != expectedName {
		t.Fatalf("Job name should be %v, was %v", expectedName, jobConfig.FullName)
	}
}

func TestCreateJob(t *testing.T) {
	name := "my_test_job"
	config := JobConfig{}
	err := jobService.CreateJob("job/Bridge%20Career/", name, &config)
	if err != nil {
		t.Fatal(err)
	}

	config.Description = "my new desc 3"
	err = jobService.UpdateJob(folder, name, &config)
	if err != nil {
		t.Fatal(err)
	}

	err = jobService.DeleteJob("job/Bridge%20Career/", name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetJobConfig(t *testing.T) {
	jobConfig, err := jobService.GetJobConfig(folder, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	if jobConfig.Description != "" {
		t.Fatalf("Job config description should be %v, was %v", "", jobConfig.Description)
	}
	if jobConfig.Definition.SCM.ConfigVersion != 2 {
		t.Fatalf("Job scm config version should be %v, was %v", 2, jobConfig.Definition.SCM.ConfigVersion)
	}
	if (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].Url != "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git" {
		t.Fatalf("Job scm url should be %v, was %v", "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git", (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].Url)
	}
	if (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].CredentialsId != "44aa91d6-ab24-498a-b2b4-911bcb17cc35" {
		t.Fatalf("Job scm CredentialsId should be %v, was %v", "44aa91d6-ab24-498a-b2b4-911bcb17cc35", (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].CredentialsId)
	}
	if (*jobConfig.Definition.SCM.Branches.Items)[0].Name != "FETCH_HEAD" {
		t.Fatalf("Job scm branch name should be %v, was %v", "FETCH_HEAD", (*jobConfig.Definition.SCM.Branches.Items)[0].Name)
	}
	if jobConfig.Definition.ScriptPath != "migrations.Jenkinsfile" {
		t.Fatalf("Job scm branch name should be %v, was %v", "migrations.Jenkinsfile", jobConfig.Definition.ScriptPath)
	}
}
