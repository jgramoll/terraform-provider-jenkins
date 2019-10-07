package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

// Job in jenkins
type job struct {
	Name     string `mapstructure:"name"`
	Plugin   string `mapstructure:"plugin"`
	Disabled bool   `mapstructure:"disabled"`

	Actions interfaceJobActions `mapstructure:"action"`
}

func newJob() *job {
	return &job{}
}

func (j *job) toClientJob(jobId string) (*client.Job, error) {
	job := client.NewJob()
	job.Id = jobId
	job.Plugin = j.Plugin
	job.Name = j.Name
	job.Disabled = j.Disabled

	actions, err := j.Actions.toClientActions()
	if err != nil {
		return nil, err
	}
	job.Actions = actions
	return job, nil
}

func JobfromClientJob(clientJob *client.Job) (*job, error) {
	j := job{}
	j.Plugin = clientJob.Plugin
	j.Name = clientJob.Name
	j.Disabled = clientJob.Disabled
	actions, err := j.Actions.fromClientActions(clientJob.Actions)
	if err != nil {
		return nil, err
	}
	j.Actions = *actions
	return &j, nil
}

func (j *job) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("plugin", j.Plugin); err != nil {
		return err
	}
	if err := d.Set("name", j.Name); err != nil {
		return err
	}
	if err := d.Set("disabled", j.Disabled); err != nil {
		return err
	}
	if err := d.Set("action", j.Actions); err != nil {
		return err
	}
	return nil
}
