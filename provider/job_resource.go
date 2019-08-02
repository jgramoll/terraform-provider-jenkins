package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/multilock"
	"github.com/mitchellh/mapstructure"
)

var jobLock = multilock.NewBasicMultiLock()

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
				ForceNew:    true,
			},
			"plugin": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name and id of the plugin",
				Optional:    true,
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
	j := newJob()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, j); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	jobId := id.String()

	jobService := m.(*Services).JobService
	err = jobService.CreateJob(j.toClientJob(jobId))
	if err != nil {
		return err
	}

	d.SetId(jobId)
	log.Println("[DEBUG] Creating job:", j.Name)
	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("name").(string)

	jobService := m.(*Services).JobService
	jobLock.RLock(jobName)
	j, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Got job %s", j.Name)
	return JobfromClientJob(j).setResourceData(d)
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	j := newJob()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &j); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(j.Name)
	err := jobService.UpdateJob(j.toClientJob(d.Id()))
	jobLock.Unlock(j.Name)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job:", d.Id())
	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("name").(string)

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	err := jobService.DeleteJob(jobName)
	jobLock.Unlock(jobName)

	log.Println("[DEBUG] Deleted job:", d.Id())
	d.SetId("")
	return err
}
