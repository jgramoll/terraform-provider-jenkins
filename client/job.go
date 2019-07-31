package client

import (
	"errors"
	"strings"
)

// ErrJobPropertyNotFound job property not found
var ErrJobPropertyNotFound = errors.New("Could not find job property")

// ErrJobActionNotFound job action not found
var ErrJobActionNotFound = errors.New("Could not find job action")

// Job
type Job struct {
	Id               string
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
		job.Id = config.Id
		job.Actions = config.Actions
		job.Disabled = config.Disabled
		job.Properties = config.Properties
		if config.Definition != nil {
			job.Definition = config.Definition.Item
		}
	}

	return job
}

func (job *Job) GetProperty(propertyId string) (JobProperty, error) {
	properties := *(job.Properties).Items
	for _, property := range properties {
		if property.GetId() == propertyId {
			return property, nil
		}
	}
	return nil, ErrJobPropertyNotFound
}

func (job *Job) DeleteProperty(propertyId string) error {
	properties := *job.Properties.Items
	for i, property := range properties {
		if property.GetId() == propertyId {
			*job.Properties.Items = append(properties[:i], properties[i+1:]...)
			return nil
		}
	}
	return ErrJobPropertyNotFound
}

func (job *Job) GetAction(actionId string) (JobAction, error) {
	for _, action := range *job.Actions.Items {
		if action.GetId() == actionId {
			return action, nil
		}
	}
	return nil, ErrJobActionNotFound
}

func (job *Job) UpdateAction(action JobAction) error {
	actions := *job.Actions.Items
	actionId := action.GetId()
	for i, oldAction := range actions {
		if oldAction.GetId() == actionId {
			actions[i] = action
			return nil
		}
	}
	return ErrJobActionNotFound
}

func (job *Job) DeleteAction(actionId string) error {
	actions := *job.Actions.Items
	for i, action := range actions {
		if action.GetId() == actionId {
			*job.Actions.Items = append(actions[:i], actions[i+1:]...)
			return nil
		}
	}
	return ErrJobActionNotFound
}
