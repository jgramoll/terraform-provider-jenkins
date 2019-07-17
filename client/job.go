package client

import (
	"errors"
	"strings"
)

// ErrJobPropertyNotFound job property not found
var ErrJobPropertyNotFound = errors.New("Could not find job property")

// Job
type Job struct {
	Id               string
	Name             string
	Disabled         bool
	Description      string
	KeepDependencies bool
	Properties       *JobProperties
	Definition       JobDefinition
}

// NewJob return Job object with default values
func NewJob() *Job {
	return &Job{
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
		job.Id = config.Id
		job.Disabled = config.Disabled
		job.Properties = config.Properties
		if config.Definition != nil {
			job.Definition = config.Definition.Item
		}
	}

	return job
}

func (job *Job) GetProperty(propertyId string) (JobProperty, error) {
	properties := *(*job.Properties).Items
	for _, property := range properties {
		if property.GetId() == propertyId {
			return property, nil
		}
	}
	return nil, ErrJobPropertyNotFound
}

func (job *Job) DeleteProperty(propertyId string) error {
	properties := *(*job.Properties).Items
	for i, property := range properties {
		if property.GetId() == propertyId {
			properties = append(properties[:i], properties[i+1:]...)
			job.Properties.Items = &properties
			return nil
		}
	}
	return ErrJobPropertyNotFound
}
