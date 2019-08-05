package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrInvalidDefinitionId
var ErrInvalidDefinitionId = errors.New("Invalid definition id, must be jobName_definitionId")

func resourceJobDefinitionParseId(input string) (jobName string, definitionId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 2 {
		err = ErrInvalidDefinitionId
		return
	}
	jobName = parts[0]
	definitionId = parts[1]
	return
}

func ResourceJobDefinitionId(jobName string, definitionId string) string {
	return joinWithIdDelimiter(jobName, definitionId)
}

func resourceJobDefinitionImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, _, err := resourceJobDefinitionParseId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("job", jobName)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobDefinitionCreate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobName := d.Get("job").(string)
	definition := createJobDefinition()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &definition); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	definitionId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Definition = definition.toClientDefinition(definitionId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobDefinitionId(jobName, definitionId))
	log.Println("[DEBUG] Creating job definition:", d.Id())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionRead(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobService := m.(*Services).JobService
	jobName, _, err := resourceJobDefinitionParseId(d.Id())
	if err != nil {
		return err
	}
	jobLock.RLock(jobName)
	j, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	definition, err := createJobDefinition().fromClientJobDefintion(j.Definition)
	if err != nil {
		return err
	}
	log.Println("[INFO] Reading from job definition", d.Id())
	return definition.setResourceData(d)
}

func resourceJobDefinitionUpdate(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobName, definitionId, err := resourceJobDefinitionParseId(d.Id())
	if err != nil {
		return err
	}

	definition := createJobDefinition()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &definition); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Definition = definition.toClientDefinition(definitionId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job definition:", d.Id())
	return resourceJobDefinitionRead(d, m, createJobDefinition)
}

func resourceJobDefinitionDelete(d *schema.ResourceData, m interface{}, createJobDefinition func() jobDefinition) error {
	jobName, _, err := resourceJobDefinitionParseId(d.Id())
	if err != nil {
		return err
	}
	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
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
