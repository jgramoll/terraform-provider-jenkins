package main

import (
	"fmt"
	"testing"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func testParametersDefinitionProperty() *client.JobParametersDefinitionProperty {
	property := client.NewJobParametersDefinitionProperty()
	testParameter1 := client.NewJobParameterDefinitionChoice()
	testParameter1.Name = "choic 1"
	testParameter1.Description = "desc"
	*testParameter1.Choices.Items.Items = []string{"1", "a", "alpha"}
	property.ParameterDefinitions =	property.ParameterDefinitions.Append(testParameter1)
	testParameter2 := client.NewJobParameterDefinitionChoice()
	testParameter2.Name = "choic 2"
	testParameter2.Description = "desc"
	*testParameter2.Choices.Items.Items = []string{"2", "b", "beta"}
	property.ParameterDefinitions =	property.ParameterDefinitions.Append(testParameter2)
	return property
}

func TestEnsureJobParametersDefinitionProperty(t *testing.T) {
	job := client.NewJob()
	property := testParametersDefinitionProperty()
	job.Properties = job.Properties.Append(property)

	if err := ensureJob(job); err != nil {
		t.Fatal(err)
	}
	if property.Id == "" {
		t.Fatalf("Did not set Parameters Definition Property Id")
	}
	definitions := *property.ParameterDefinitions.Items
	if definitions[0].GetId() == "" {
		t.Fatalf("Did not set Parameter Definition 1 Id")
	}
	if definitions[1].GetId() == "" {
		t.Fatalf("Did not set Parameter Definition 2 Id")
	}
}

func TestJobParametersDefinitionPropertyCode(t *testing.T) {
	job := client.NewJob()
	property := testParametersDefinitionProperty()
	job.Properties = job.Properties.Append(property)

	result := jobCode(job)
	expected := `resource "jenkins_job" "main" {
	name     = ""
	plugin   = ""
	disabled = false
}

resource "jenkins_job_parameters_definition_property" "property_1" {
	job = "${jenkins_job.main.name}"
}

resource "jenkins_job_parameter_definition_choice" "parameter_1_1" {
	property = "${jenkins_job_parameters_definition_property.property_1.id}"

	name        = "choic 1"
	description = "desc"
	choices     = ["1", "a", "alpha"]
}

resource "jenkins_job_parameter_definition_choice" "parameter_1_2" {
	property = "${jenkins_job_parameters_definition_property.property_1.id}"

	name        = "choic 2"
	description = "desc"
	choices     = ["2", "b", "beta"]
}
`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform code not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestJobParametersDefinitionPropertyImportScript(t *testing.T) {
	job := client.NewJob()
	job.Id = "id"
	job.Name = "name"
	property := testParametersDefinitionProperty()
	property.Id = "paramPropertyId"
	(*property.ParameterDefinitions.Items)[0].SetId("c1")
	(*property.ParameterDefinitions.Items)[1].SetId("c2")
	job.Properties = job.Properties.Append(property)

	result := jobImportScript(job)
	expected := fmt.Sprintf(`terraform init

terraform import jenkins_job.main "name"

terraform import jenkins_job_parameters_definition_property.property_1 "%v"

terraform import jenkins_job_parameter_definition_choice.parameter_1_1 "%v"

terraform import jenkins_job_parameter_definition_choice.parameter_1_2 "%v"
`, provider.ResourceJobPropertyId(job.Name, property.Id),
provider.ResourceJobParameterDefinitionId(job.Name, property.Id, "c1"),
provider.ResourceJobParameterDefinitionId(job.Name, property.Id, "c2"))

	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, false)
		t.Fatalf("job terraform import script not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}
