package client

import (
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

func TestGetJob(t *testing.T) {
	jobName := "Bridge Career/migrations_change"
	jobConfig, err := jobService.GetJob(jobName)
	if err != nil {
		t.Fatal(err)
	}
	if jobConfig.Name != "Bridge Career/migrations_change" {
		t.Fatalf("Job name should be %v, was %v", "Bridge Career/migrations_change", jobConfig.Name)
	}
	if jobConfig.Description != "" {
		t.Fatalf("Job config description should be %v, was %v", "", jobConfig.Description)
	}
	// if jobConfig.Definition.SCM.ConfigVersion != 2 {
	// 	t.Fatalf("Job scm config version should be %v, was %v", 2, jobConfig.Definition.SCM.ConfigVersion)
	// }
	// if (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].Url != "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git" {
	// 	t.Fatalf("Job scm url should be %v, was %v", "ssh://gerrit.instructure.com:29418/bridge-career-infrastructure.git", (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].Url)
	// }
	// if (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].CredentialsId != "44aa91d6-ab24-498a-b2b4-911bcb17cc35" {
	// 	t.Fatalf("Job scm CredentialsId should be %v, was %v", "44aa91d6-ab24-498a-b2b4-911bcb17cc35", (*jobConfig.Definition.SCM.UserRemoteConfigs.Items)[0].CredentialsId)
	// }
	// if (*jobConfig.Definition.SCM.Branches.Items)[0].Name != "FETCH_HEAD" {
	// 	t.Fatalf("Job scm branch name should be %v, was %v", "FETCH_HEAD", (*jobConfig.Definition.SCM.Branches.Items)[0].Name)
	// }
	// if jobConfig.Definition.ScriptPath != "migrations.Jenkinsfile" {
	// 	t.Fatalf("Job scm branch name should be %v, was %v", "migrations.Jenkinsfile", jobConfig.Definition.ScriptPath)
	// }
}

func TestCreateJob(t *testing.T) {
	job := Job{Name: "Bridge Career/my_test_job"}
	jobName := job.Name

	err := jobService.CreateJob(&job)
	if err != nil {
		t.Fatal(err)
	}

	job.Description = "my new desc 3"
	err = jobService.UpdateJob(&job)
	if err != nil {
		t.Fatal(err)
	}

	err = jobService.DeleteJob(jobName)
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
