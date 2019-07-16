package provider

// type Comparable struct {
// 	CompareType string `mapstructure:"compare_type"`
// 	Pattern     string `mapstructure:"pattern"`
// }

// // Job in jenkins
// type jobGerritTrigger struct {
// 	Project *Comparable `mapstructure:"project"`
// 	Branch  *Comparable `mapstructure:"branch"`
// 	// Disabled bool   `mapstructure:"disabled"`
// }

// func newJobGerritTrigger() *jobGerritTrigger {
// 	return &jobGerritTrigger{}
// }

// // func (j *job) toClientJob() *client.Job {
// // 	return &client.Job{
// // 		Name:     j.Name,
// // 		Disabled: j.Disabled,
// // 	}
// // }

// // func JobfromClientJob(j *client.Job) *job {
// // 	return &job{
// // 		Name:     j.Name,
// // 		Disabled: j.Disabled,
// // 	}
// // }

// // func (j *job) setResourceData(d *schema.ResourceData) error {
// // 	d.SetId(j.Name)
// // 	err := d.Set("name", j.Name)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	err = d.Set("disabled", j.Disabled)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// // // JobFromResourceData get job from resource data
// // func JobFromResourceData(job *client.Job, d *schema.ResourceData) {
// // 	job.Name = d.Get("name").(string)
// // 	job.Disabled = d.Get("disabled").(bool)
// // }
