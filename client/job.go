package client

import (
	"strings"
)

// Job
type Job struct {
	Plugin           string
	Name             string
	Disabled         bool
	Actions          *JobActions `xml:"actions"`
	Description      string
	KeepDependencies bool
	Properties       *JobProperties
	Definition       JobDefinition
}

// NewJob return Job object with default values
func NewJob() *Job {
	return &Job{
		Actions:          NewJobActions(),
		KeepDependencies: false,
		Properties:       NewJobProperties(),
	}
}

func (job *Job) Folder() string {
	nameParts := strings.Split(job.Name, "/")
	return strings.Join(nameParts[:len(nameParts)-1], "/")
}

func (job *Job) NameOnly() string {
	nameParts := strings.Split(job.Name, "/")
	return nameParts[len(nameParts)-1]
}

func newJobFromConfigAndDetails(config *jobConfig, details *jobDetails) *Job {
	job := NewJob()

	if details != nil {
		job.Name = details.FullName
		job.Description = details.Description
	}

	if config != nil {
		job.Plugin = config.Plugin
		job.Actions = config.Actions
		job.Disabled = config.Disabled
		job.Properties = config.Properties
		if config.Definition != nil {
			job.Definition = config.Definition.Item
		}
	}

	return job
}
