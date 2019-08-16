package client

import (
	"encoding/xml"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var paramterDefinitionChoice *JobParameterDefinitionChoice

func init() {
	paramterDefinitionChoice = NewJobParameterDefinitionChoice()
	paramterDefinitionChoice.Name = "ENV"
	paramterDefinitionChoice.Choices = paramterDefinitionChoice.Choices.Append("edge")
	paramterDefinitionChoice.Choices = paramterDefinitionChoice.Choices.Append("staging")
	paramterDefinitionChoice.Choices = paramterDefinitionChoice.Choices.Append("prod")
}

func TestJobParametersDefinitionPropertyDeserialize(t *testing.T) {
	var job Job
	err := xml.Unmarshal([]byte(expectedParameterDefinitionChoiceJson), &job)
	if err != nil {
		t.Fatalf("failed to deserialize xml %s", err)
	}
	if job.Properties == nil || len(*job.Properties.Items) == 0 {
		t.Fatalf("failed to deserialize properties")
	}
	property, ok := (*job.Properties.Items)[0].(*JobParametersDefinitionProperty)
	if !ok {
		t.Fatalf("failed to deserialize JobParametersDefinitionProperty")
	}
	if property.ParameterDefinitions == nil || property.ParameterDefinitions.Items == nil || len(*property.ParameterDefinitions.Items) == 0 {
		t.Fatalf("failed to deserialize property items")
	}
	resultChoiceProperty, ok := (*property.ParameterDefinitions.Items)[0].(*JobParameterDefinitionChoice)
	if !ok {
		t.Fatalf("failed to deserialize JobParameterDefinitionChoice")
	}
	if resultChoiceProperty.Name != paramterDefinitionChoice.Name {
		t.Fatalf("failed to deserialize JobParameterDefinitionChoice Name")
	}
}

func TestJobParametersDefinitionPropertySerialize(t *testing.T) {
	job := NewJob()
	jobParametersDefinitionProperty := NewJobParametersDefinitionProperty()
	jobParametersDefinitionProperty.ParameterDefinitions =
		jobParametersDefinitionProperty.ParameterDefinitions.Append(paramterDefinitionChoice)
	job.Properties = job.Properties.Append(jobParametersDefinitionProperty)

	config := JobConfigFromJob(job)
	resultBytes, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)

	if result != expectedParameterDefinitionChoiceJson {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedParameterDefinitionChoiceJson, result, true)
		t.Fatalf("job definition not expected: %s", dmp.DiffPrettyText(diffs))
	}
}

var expectedParameterDefinitionChoiceJson = `<flow-definition>
	<actions></actions>
	<description></description>
	<keepDependencies>false</keepDependencies>
	<properties>
		<hudson.model.ParametersDefinitionProperty>
			<parameterDefinitions>
				<hudson.model.ChoiceParameterDefinition>
					<name>ENV</name>
					<description></description>
					<choices class="java.util.Arrays$ArrayList">
						<a class="string-array">
							<string>edge</string>
							<string>staging</string>
							<string>prod</string>
						</a>
					</choices>
				</hudson.model.ChoiceParameterDefinition>
			</parameterDefinitions>
		</hudson.model.ParametersDefinitionProperty>
	</properties>
	<triggers></triggers>
	<disabled>false</disabled>
</flow-definition>`
