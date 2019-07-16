package client

import (
	"encoding/xml"
	"testing"
)

func TestJobConfigSerialize(t *testing.T) {
	job := NewJob()
	config := JobConfigFromJob(job)
	resultBytes, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)
	expected := `<flow-definition id="" plugin="workflow-job@2.33">
	<description></description>
	<properties></properties>
	<disabled>false</disabled>
</flow-definition>`
	if result != expected {
		t.Fatalf("job definition should be %s, was %s", expected, result)
	}
}
