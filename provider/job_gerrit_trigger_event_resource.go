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

// ErrInvalidJobTriggerEventId
var ErrInvalidJobTriggerEventId = errors.New("Invalid trigger event id, must be jobName_propertyId_triggerId_eventId")

func resourceJobTriggerEventId(input string) (jobName string, propertyId string, triggerId string, eventId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 4 {
		err = ErrInvalidJobTriggerEventId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	triggerId = parts[2]
	eventId = parts[3]
	return
}

func resourceJobTriggerEventImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, triggerId, _, err := resourceJobTriggerEventId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("trigger", strings.Join([]string{jobName, propertyId, triggerId}, IdDelimiter))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobTriggerEventCreate(d *schema.ResourceData, m interface{}, createJobTriggerEvent func() jobGerritTriggerEvent) error {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(d.Get("trigger").(string))
	if err != nil {
		return err
	}

	event := createJobTriggerEvent()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &event); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	eventId := id.String()

	clientEvent, err := event.toClientJobTriggerEvent(eventId)
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

	p, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property := p.(*client.JobPipelineTriggersProperty)
	trigger, err := property.GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	gerritTrigger := trigger.(*client.JobGerritTrigger)
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(clientEvent)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, propertyId, triggerId, eventId}, IdDelimiter))
	log.Println("[DEBUG] Creating job trigger:", d.Id())
	return resourceJobTriggerEventRead(d, m, createJobTriggerEvent)
}

func resourceJobTriggerEventRead(d *schema.ResourceData, m interface{}, createJobTriggerEvent func() jobGerritTriggerEvent) error {
	jobName, propertyId, triggerId, eventId, err := resourceJobTriggerEventId(d.Id())
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
		log.Println("[WARN] No Property found:", err)
		d.SetId("")
		return err
	}
	property := clientProperty.(*client.JobPipelineTriggersProperty)
	trigger, err := property.GetTrigger(triggerId)
	if err != nil {
		return err
	}
	gerritTrigger := trigger.(*client.JobGerritTrigger)
	clientEvent, err := gerritTrigger.GetEvent(eventId)
	if err != nil {
		return err
	}
	event, err := createJobTriggerEvent().fromClientJobTriggerEvent(clientEvent)
	if err != nil {
		return err
	}
	log.Println("[INFO] Updating state for job trigger event", d.Id())
	return event.setResourceData(d)
}

func resourceJobTriggerEventUpdate(d *schema.ResourceData, m interface{}, createJobTriggerEvent func() jobGerritTriggerEvent) error {
	jobName, propertyId, triggerId, eventId, err := resourceJobTriggerEventId(d.Id())
	if err != nil {
		return err
	}

	event := createJobTriggerEvent()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &event); err != nil {
		return err
	}
	clientEvent, err := event.toClientJobTriggerEvent(eventId)
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

	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	trigger, err := property.GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	gerritTrigger := trigger.(*client.JobGerritTrigger)
	err = gerritTrigger.UpdateEvent(clientEvent)
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
	return resourceJobTriggerEventRead(d, m, createJobTriggerEvent)
}

func resourceJobTriggerEventDelete(d *schema.ResourceData, m interface{}, createJobTriggerEvent func() jobGerritTriggerEvent) error {
	jobName, propertyId, triggerId, eventId, err := resourceJobTriggerEventId(d.Id())
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

	propertyInterface, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	property := propertyInterface.(*client.JobPipelineTriggersProperty)
	trigger, err := property.GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	gerritTrigger := trigger.(*client.JobGerritTrigger)
	err = gerritTrigger.DeleteEvent(eventId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
