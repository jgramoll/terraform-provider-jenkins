package client

import (
	"encoding/xml"
	"testing"
)

func TestCpsScmFlowDefinitionSerialize(t *testing.T) {
	definition := NewCpsScmFlowDefinition()
	resultBytes, err := xml.MarshalIndent(definition, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)
	expected := `<CpsScmFlowDefinition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition">
	<scriptPath></scriptPath>
	<lightweight>false</lightweight>
</CpsScmFlowDefinition>`
	if result != expected {
		t.Fatalf("job definition should be %s, was %s", expected, result)
	}
}
