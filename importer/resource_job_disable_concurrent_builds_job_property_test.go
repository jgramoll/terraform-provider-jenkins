package main

import (
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testDisableConcurrentBuildsJobProperty() *client.JobDisableConcurrentBuildsJobProperty {
	property := client.NewJobDisableConcurrentBuildsJobProperty()
	return property
}

func TestJobDisableConcurrentBuildsJobPropertyCode(t *testing.T) {
	job := client.NewJob()
	property := testDisableConcurrentBuildsJobProperty()
	job.Properties = job.Properties.Append(property)

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
	name     = ""
	plugin   = ""
	disabled = false
}

resource "jenkins_job_disable_concurrent_builds_property" "property_1" {
	job = "${jenkins_job.main.name}"
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
