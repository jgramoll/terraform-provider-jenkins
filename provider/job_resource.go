package provider

import (
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	// "github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// ErrMissingJobName missing job name
var ErrMissingJobName = errors.New("job name must be provided")

func jobResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job, including folder heirarchy. E.g. Foo/Bar/Baz",
				Required:    true,
			},
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the job is disabled",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {
	var j job
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &j); err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job:", j.Name)
	jobService := m.(*Services).JobService
	err := jobService.CreateJob(j.toClientJob())
	if err != nil {
		return err
	}

	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Id())
	if err != nil {
		log.Println("[WARN] No Job found:", d.Id())
		d.SetId("")
		return err
	}

	log.Printf("[INFO] Got job %s", j.Name)
	return JobfromClientJob(j).setResourceData(d)
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Id())
	if err != nil {
		return err
	}
	JobFromResourceData(j, d)

	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job:", d.Id())
	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Id()
	log.Println("[DEBUG] Deleting job:", d.Id())
	d.SetId("")
	jobService := m.(*Services).JobService
	return jobService.DeleteJob(name)
}
