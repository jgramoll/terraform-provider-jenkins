package provider

import (
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

var jobLock sync.Mutex

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

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	j.RefId = id.String()

	log.Println("[DEBUG] Creating job:", j.Name)
	jobService := m.(*Services).JobService
	err = jobService.CreateJob(j.toClientJob())
	if err != nil {
		return err
	}

	d.SetId(j.RefId)
	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("name").(string))
	if err != nil {
		log.Println("[WARN] No Job found:", d.Id())
		d.SetId("")
		return err
	}

	log.Printf("[INFO] Got job %s", j.Name)
	return JobfromClientJob(j).setResourceData(d)
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("name").(string))
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
	jobLock.Lock()
	defer jobLock.Unlock()

	name := d.Id()
	log.Println("[DEBUG] Deleting job:", d.Id())
	d.SetId("")
	jobService := m.(*Services).JobService
	return jobService.DeleteJob(name)
}
