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

	Actions    *interfaceJobActions    `mapstructure:"action"`
	Definition *interfaceJobDefinition `mapstructure:"definition"`
	Properties *interfaceJobProperties `mapstructure:"property"`
}

func newJob() *job {
	return &job{}
}

func (j *job) toClientJob() (*client.Job, error) {
	job := client.NewJob()
	job.Plugin = j.Plugin
	job.Name = j.Name
	job.Disabled = j.Disabled

	actions, err := j.Actions.toClientActions()
	if err != nil {
		return nil, err
	}
	job.Actions = actions

	definition, err := j.Definition.toClientDefinition()
	if err != nil {
		return nil, err
	}
	job.Definition = definition

	properties, err := j.Properties.toClientProperties()
	if err != nil {
		return nil, err
	}
	job.Properties = properties

	return job, nil
}

func JobfromClientJob(clientJob *client.Job) (*job, error) {
	j := newJob()
	j.Plugin = clientJob.Plugin
	j.Name = clientJob.Name
	j.Disabled = clientJob.Disabled

	actions, err := j.Actions.fromClientActions(clientJob.Actions)
	if err != nil {
		return nil, err
	}
	j.Actions = actions

	definition, err := j.Definition.fromClientDefinition(clientJob.Definition)
	if err != nil {
		return nil, err
	}
	j.Definition = definition

	properties, err := j.Properties.fromClientProperties(clientJob.Properties)
	if err != nil {
		return nil, err
	}
	j.Properties = properties

	return j, nil
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
	if err := d.Set("definition", j.Definition); err != nil {
		return err
	}
	if err := d.Set("property", j.Properties); err != nil {
		return err
	}
	return nil
}
