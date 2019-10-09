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
