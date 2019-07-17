package client

import (
	"testing"
)

func TestJobDetailNameOnly(t *testing.T) {
	job := NewJob()
	job.Name = "fee fi/fo/fum"
	expected := "fum"
	if job.NameOnly() != expected {
		t.Fatalf("job name only should be %v, was %v", expected, job.NameOnly())
	}
}

func TestJobDetailFolder(t *testing.T) {
	job := NewJob()
	job.Name = "fee fi/fo/fum"
	expected := "fee fi/fo"
	if job.Folder() != expected {
		t.Fatalf("job folder should be %v, was %v", expected, job.Folder())
	}
}

func TestJobGetProperty(t *testing.T) {
	job := NewJob()
	property1 := NewJobPipelineTriggersProperty()
	property1.Id = "1"
	property2 := NewJobPipelineBuildDiscarderProperty()
	property2.Id = "2"
	job.Properties = job.Properties.Append(property1)
	job.Properties = job.Properties.Append(property2)

	result, err := job.GetProperty("1")
	if err != nil {
		t.Fatal(err)
	}
	if result != property1 {
		t.Fatalf("job get property should return %v, was %v", property1, result)
	}
}

func TestJobDeleteProperty(t *testing.T) {
	job := NewJob()
	property1 := NewJobPipelineTriggersProperty()
	property1.Id = "1"
	property2 := NewJobPipelineBuildDiscarderProperty()
	property2.Id = "2"
	job.Properties = job.Properties.Append(property1)
	job.Properties = job.Properties.Append(property2)

	err := job.DeleteProperty("1")
	if err != nil {
		t.Fatal(err)
	}

	_, err = job.GetProperty("1")
	if err != ErrJobPropertyNotFound {
		t.Fatal("Should fail to find property")
	}
}
