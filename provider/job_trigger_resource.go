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

// ErrInvalidJobTriggerId
var ErrInvalidJobTriggerId = errors.New("Invalid trigger id, must be jobName_propertyId_triggerId")

func resourceJobTriggerId(input string) (jobName string, propertyId string, triggerId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidJobTriggerId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	triggerId = parts[2]
	return
}

func ResourceJobTriggerId(jobName string, propertyId string, triggerId string) string {
	return joinWithIdDelimiter(jobName, propertyId, triggerId)
}

func resourceJobTriggerImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, _, err := resourceJobTriggerId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("property", ResourceJobPropertyId(jobName, propertyId))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobTriggerCreate(d *schema.ResourceData, m interface{}, createJobTrigger func() jobTrigger) error {
	jobName, propertyId, err := resourceJobPropertyParseId(d.Get("property").(string))
	if err != nil {
		return err
	}

	trigger := createJobTrigger()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &trigger); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	triggerId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	p, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property := p.(*client.JobPipelineTriggersProperty)
	clientTrigger, err := trigger.toClientJobTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property.Triggers = property.Triggers.Append(clientTrigger)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobTriggerId(jobName, propertyId, triggerId))
	log.Println("[DEBUG] Creating job trigger:", d.Id())
	return resourceJobTriggerRead(d, m, createJobTrigger)
}

func resourceJobTriggerRead(d *schema.ResourceData, m interface{}, createJobTrigger func() jobTrigger) error {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(d.Id())
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
	clientProperty, err := j.GetProperty(propertyId)
	if err != nil {
		log.Println("[WARN] No Job Property found:", err)
		d.SetId("")
		return nil
	}
	property := clientProperty.(*client.JobPipelineTriggersProperty)
	clientTrigger, err := property.GetTrigger(triggerId)
	if err != nil {
		log.Println("[WARN] No Job Trigger found:", err)
		d.SetId("")
		return nil
	}
	trigger, err := createJobTrigger().fromClientJobTrigger(clientTrigger)
	if err != nil {
		return err
	}
	log.Println("[INFO] Updating state for job trigger", d.Id())
	return trigger.setResourceData(d)
}

func resourceJobTriggerUpdate(d *schema.ResourceData, m interface{}, createJobTrigger func() jobTrigger) error {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(d.Id())
	if err != nil {
		return err
	}

	trigger := createJobTrigger()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &trigger); err != nil {
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
		return err
	}
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	clientTrigger, err := trigger.toClientJobTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = property.UpdateTrigger(clientTrigger)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updating job trigger", d.Id())
	return resourceJobTriggerRead(d, m, createJobTrigger)
}

func resourceJobTriggerDelete(d *schema.ResourceData, m interface{}, createJobTrigger func() jobTrigger) error {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(d.Id())
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Job Trigger:", err)
		d.SetId("")
		return nil
	}

	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Job Trigger:", err)
		d.SetId("")
		return nil
	}
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	err = property.DeleteTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Job Trigger:", err)
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
