package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceJobDefinitionCreate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	s := createJobDefinition()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	definition := s.(jobDefinition)

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	definition.setRefID(id.String())

	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Definition = definition.toClientDefinition()
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job definition:", id)
	d.SetId(id.String())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionUpdate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobName := d.Get("job").(string)

	s := createJobDefinition()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	definition := s.(jobDefinition)
	definition.setRefID(d.Id())

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Definition = definition.toClientDefinition()
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job definition:", d.Id())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionRead(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.RLock(jobName)
	j, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	definition := createJobDefinition()
	definition = definition.fromClientJobDefintion(j.Definition)
	if definition == nil {
		return nil
	}

	log.Println("[INFO] Updating from job definition", definition)
	return definition.setResourceData(d)
}

func resourceJobDefinitionDelete(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {

	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Definition = nil
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
