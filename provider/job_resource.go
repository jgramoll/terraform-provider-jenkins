package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

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
			"action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     jobActionResource(),
			},
			"definition": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     jobDefinitionResource(),
			},
			"property": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     jobPropertyResource(),
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

	clientJob, err := j.toClientJob()
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	err = jobService.CreateJob(clientJob)
	if err != nil {
		return err
	}

	d.SetId(j.Name)
	log.Println("[DEBUG] Creating job:", j.Name)
	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	jobName := d.Id()

	jobService := m.(*Services).JobService
	clientJob, err := jobService.GetJob(jobName)
	if err != nil {
		println("WHAT!", err.Error())
		log.Println("[WARN] No Job found:", d.Id())
		d.SetId("")
		return nil
	}

	j, err := JobfromClientJob(clientJob)
	if err != nil {
		println("WHAT2!")
		log.Println("[WARN] Invalid Job:", d.Id())
		return nil
	}
	log.Printf("[INFO] Got job %s", j.Name)
	return j.setResourceData(d)
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	j := newJob()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &j); err != nil {
		return err
	}

	clientJob, err := j.toClientJob()
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	err = jobService.UpdateJob(clientJob)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job:", d.Id())
	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	jobName := d.Id()
	jobService := m.(*Services).JobService
	err := jobService.DeleteJob(jobName)

	log.Println("[DEBUG] Deleted job:", d.Id())
	d.SetId("")
	return err
}
