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
}

func newJob() *job {
	return &job{}
}

func (j *job) toClientJob(jobId string) *client.Job {
	job := client.NewJob()
	job.Id = jobId
	job.Plugin = j.Plugin
	job.Name = j.Name
	job.Disabled = j.Disabled
	return job
}

func JobfromClientJob(clientJob *client.Job) *job {
	j := job{}
	j.Plugin = clientJob.Plugin
	j.Name = clientJob.Name
	j.Disabled = clientJob.Disabled
	return &j
}

func (j *job) setResourceData(d *schema.ResourceData) error {
	d.SetId(j.Name)
	if err := d.Set("plugin", j.Plugin); err != nil {
		return err
	}
	if err := d.Set("name", j.Name); err != nil {
		return err
	}
	if err := d.Set("disabled", j.Disabled); err != nil {
		return err
	}
	return nil
}
