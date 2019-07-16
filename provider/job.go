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

func (j *job) toClientJob() *client.Job {
	job := client.NewJob()
	// Todo pass in id / data
	// job.Id = j.RefId
	job.Name = j.Name
	job.Disabled = j.Disabled
	return job
}

func JobfromClientJob(j *client.Job) *job {
	return &job{
		Name:     j.Name,
		Disabled: j.Disabled,
	}
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

// TODO can we get rid of this?
// JobFromResourceData get job from resource data
func JobFromResourceData(job *client.Job, d *schema.ResourceData) {
	job.Name = d.Get("name").(string)
	job.Disabled = d.Get("disabled").(bool)
}
