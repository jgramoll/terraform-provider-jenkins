package client

import (
	"encoding/xml"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var disableConcurrentBuildsJobProperty *JobDisableConcurrentBuildsJobProperty

func init() {
	disableConcurrentBuildsJobProperty = NewJobDisableConcurrentBuildsJobProperty()
}

func TestJobDisableConcurrentBuildsJobPropertyDeserialize(t *testing.T) {
	var job Job
	err := xml.Unmarshal([]byte(expectedDisableConcurrentBuildsJobPropertyJson), &job)
	if err != nil {
		t.Fatalf("failed to deserialize xml %s", err)
	}
	if job.Properties == nil || len(*job.Properties.Items) == 0 {
		t.Fatalf("failed to deserialize properties")
	}
	_, ok := (*job.Properties.Items)[0].(*JobDisableConcurrentBuildsJobProperty)
	if !ok {
		t.Fatalf("failed to deserialize JobDisableConcurrentBuildsJobProperty")
	}
}

func TestJobDisableConcurrentBuildsJobPropertySerialize(t *testing.T) {
	job := NewJob()
	jobDisableConcurrentBuildsJobProperty := NewJobDisableConcurrentBuildsJobProperty()
	job.Properties = job.Properties.Append(jobDisableConcurrentBuildsJobProperty)

	config := JobConfigFromJob(job)
	resultBytes, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)

	if result != expectedDisableConcurrentBuildsJobPropertyJson {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedDisableConcurrentBuildsJobPropertyJson, result, true)
		t.Fatalf("job definition not expected: %s", dmp.DiffPrettyText(diffs))
	}
}

var expectedDisableConcurrentBuildsJobPropertyJson = `<flow-definition>
	<actions></actions>
	<description></description>
	<keepDependencies>false</keepDependencies>
	<properties>
		<org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty></org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty>
	</properties>
	<triggers></triggers>
	<disabled>false</disabled>
</flow-definition>`
