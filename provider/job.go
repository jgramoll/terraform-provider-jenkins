package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

// Job in jenkins
type job struct {
	Name     string `mapstructure:"name"`
	Disabled bool   `mapstructure:"disabled"`
}

func newJob() *job {
	return &job{}
}

func (j *job) toClientJob(jobId string) *client.Job {
	job := client.NewJob()
	job.Id = jobId
	job.Name = j.Name
	job.Disabled = j.Disabled
	return job
}

func JobfromClientJob(clientJob *client.Job) *job {
	j := job{}
	j.Name = clientJob.Name
	j.Disabled = clientJob.Disabled
	return &j
}

func (j *job) setResourceData(d *schema.ResourceData) error {
	d.SetId(j.Name)
	err := d.Set("name", j.Name)
	if err != nil {
		return err
	}
	err = d.Set("disabled", j.Disabled)
	if err != nil {
		return err
	}
	return nil
}
