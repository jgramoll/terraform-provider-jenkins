package client

import (
	"strings"
)

// Job
type Job struct {
	Name        string
	Disabled    bool
	Description string
}

// NewJob return Job object with default values
func NewJob() *Job {
	return &Job{}
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
	job := Job{}

	if details != nil {
		job.Name = details.FullName
		job.Description = details.Description
	}

	if config != nil {
		job.Disabled = config.Disabled
	}

	return &job
}
