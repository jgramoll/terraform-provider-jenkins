package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// ErrInvalidParameterDefinitionId
var ErrInvalidParameterDefinitionId = errors.New("Invalid parameter id, must be jobName_propertyId_definitionId")

func resourceJobParameterDefinitionParseId(input string) (jobName string, propertyId string, definitionId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidParameterDefinitionId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	definitionId = parts[2]
	return
}

func ResourceJobParameterDefinitionId(jobName string, propertyId string, definitionId string) string {
	return joinWithIdDelimiter(jobName, propertyId, definitionId)
}

func resourceJobParameterDefinitionImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, _, err := resourceJobParameterDefinitionParseId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("property", ResourceJobPropertyId(jobName, propertyId))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobParameterDefinitionCreate(
	d *schema.ResourceData, m interface{}, createJobParameterDefinition func() jobParameterDefinition,
) error {
	jobName, propertyId, err := resourceJobPropertyParseId(d.Get("property").(string))
	if err != nil {
		return err
	}

	definition := createJobParameterDefinition()
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
	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	clientDefinition := definition.toClientJobParameterDefinition(definitionId)
	property.ParameterDefinitions = property.ParameterDefinitions.Append(clientDefinition)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobParameterDefinitionId(jobName, propertyId, definitionId))
	log.Println("[DEBUG] Creating job property:", d.Id())
	return resourceJobParameterDefinitionRead(d, m, createJobParameterDefinition)
}

func resourceJobParameterDefinitionRead(
	d *schema.ResourceData, m interface{}, createJobParameterDefinition func() jobParameterDefinition,
) error {
	jobName, propertyId, definitionId, err := resourceJobParameterDefinitionParseId(d.Id())
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.RLock(jobName)
	j, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}
	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		log.Println("[WARN] No Property found.", propertyId, err)
		d.SetId("")
		return nil
	}
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	clientDefinition, err := property.GetParameterDefinition(definitionId)
	if err != nil {
		log.Println("[WARN] No Parameter Definition found:", err)
		d.SetId("")
		return nil
	}

	definition, err := createJobParameterDefinition().fromClientJobParameterDefintion(clientDefinition)
	if err != nil {
		log.Println("[WARN] Invalid Property found:", err)
		return err
	}
	log.Println("[INFO] Updating from job parameter definition", d.Id())
	return definition.setResourceData(d)
}

func resourceJobParameterDefinitionUpdate(
	d *schema.ResourceData, m interface{}, createJobParameterDefinition func() jobParameterDefinition,
) error {
	jobName, propertyId, definitionId, err := resourceJobParameterDefinitionParseId(d.Id())
	if err != nil {
		return err
	}

	definition := createJobParameterDefinition()
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
	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] No Property found:", err)
		d.SetId("")
		return nil
	}
	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	clientDefinition := definition.toClientJobParameterDefinition(definitionId)
	err = property.UpdateParameterDefinition(clientDefinition)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] No Parameter Definition found:", err)
		d.SetId("")
		return nil
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updating job parameter definition", d.Id())
	return resourceJobParameterDefinitionRead(d, m, createJobParameterDefinition)
}

func resourceJobParameterDefinitionDelete(
	d *schema.ResourceData, m interface{}, createJobParameterDefinition func() jobParameterDefinition,
) error {
	jobName, propertyId, definitionId, err := resourceJobParameterDefinitionParseId(d.Id())
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could find Job:", err)
		d.SetId("")
		return nil
	}
	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] No Property found:", err)
		d.SetId("")
		return nil
	}

	property := propertyInterface.(*client.JobParametersDefinitionProperty)
	err = property.DeleteParameterDefinition(definitionId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] No Parameter Definition found:", err)
		d.SetId("")
		return nil
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
