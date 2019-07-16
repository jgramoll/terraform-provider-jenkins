package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceJobPropertyCreate(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	s := createJobProperty()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	property := s.(jobProperty)

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	property.setRefID(id.String())

	jobService := m.(*Services).JobService
	job, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	properties := append(*(*job.Properties).Items, property.toClientProperty())
	job.Properties.Items = &properties

	err = jobService.UpdateJob(job)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job property:", id)
	d.SetId(id.String())
	return resourceJobPropertyRead(d, m, createJobProperty)
}

func resourceJobPropertyRead(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobId := d.Get("job").(string)
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(jobId)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	clientProperty, err := j.GetProperty(d.Id())
	if err != nil {
		log.Println("[WARN] No Job Property found:", err)
		d.SetId("")
	} else {
		property := createJobProperty().(jobProperty)
		property = property.fromClientJobProperty(clientProperty)
		log.Println("[INFO] Updating from job property", clientProperty)
		err = property.setResourceData(d)
	}

	return err
}

func resourceJobPropertyDelete(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	p := createJobProperty()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &p); err != nil {
		return err
	}
	property := p.(jobProperty)
	property.setRefID(d.Id())

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	err = j.DeleteProperty(d.Id())
	if err != nil {
		return err
	}

	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
