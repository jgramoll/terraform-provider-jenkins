package main

import (
	"fmt"
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testDisableConcurrentBuildsJobProperty() *client.JobDisableConcurrentBuildsJobProperty {
	property := client.NewJobDisableConcurrentBuildsJobProperty()
	return property
}

func TestEnsureJobDisableConcurrentBuildsJobProperty(t *testing.T) {
	job := client.NewJob()
	property := testDisableConcurrentBuildsJobProperty()
	job.Properties = job.Properties.Append(property)

	if err := ensureJob(job); err != nil {
		t.Fatal(err)
	}
	if property.Id == "" {
		t.Fatalf("Did not set Disable Concurrent Builds Property Id")
	}
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

func TestJobDisableConcurrentBuildsJobPropertyImportScript(t *testing.T) {
	job := client.NewJob()
	job.Id = "id"
	job.Name = "name"
	property := testDisableConcurrentBuildsJobProperty()
	property.Id = "paramPropertyId"
	job.Properties = job.Properties.Append(property)

	result := jobImportScript(job)
	expected := fmt.Sprintf(`terraform init

terraform import jenkins_job.main "name"

terraform import jenkins_job_disable_concurrent_builds_property.property_1 "%v"
`, provider.ResourceJobPropertyId(job.Name, property.Id))

	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform import script not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
