package client

import (
	"encoding/xml"
	"testing"
)

func TestJobDefinitionSerialize(t *testing.T) {
	definition := &JobDefinitionXml{
		Item: NewCpsScmFlowDefinition(),
	}
	resultBytes, err := xml.MarshalIndent(definition, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)
	expected := `<JobDefinitionXml class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" id="">
	<scriptPath></scriptPath>
	<lightweight>false</lightweight>
</JobDefinitionXml>`
	if result != expected {
		t.Fatalf("job definition should be %s, was %s", expected, result)
	}
}
