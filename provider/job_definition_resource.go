package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceJobDefinitionCreate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobLock.Lock()
	defer jobLock.Unlock()

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

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	j.Definition = definition.toClientDefinition()
	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job definition:", id)
	d.SetId(id.String())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionUpdate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	s := createJobDefinition()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	definition := s.(jobDefinition)
	definition.setRefID(d.Id())

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	j.Definition = definition.toClientDefinition()
	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job definition:", d.Id())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionRead(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobId := d.Get("job").(string)
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(jobId)
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
	jobLock.Lock()
	defer jobLock.Unlock()

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	j.Definition = nil
	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
