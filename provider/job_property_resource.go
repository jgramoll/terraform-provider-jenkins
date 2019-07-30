package provider

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrInvalidPropertyId
var ErrInvalidPropertyId = errors.New("Invalid property id, must be jobName_propertyId")

func resourceJobPropertyId(propertyString string) (jobName string, propertyId string, err error) {
	parts := strings.Split(propertyString, "_")
	if len(parts) != 2 {
		err = ErrInvalidPropertyId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	return
}

func resourceJobPropertyCreate(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobName := d.Get("job").(string)
	property := createJobProperty()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &property); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	propertyId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	j.Properties = j.Properties.Append(property.toClientProperty(propertyId))
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", jobName, propertyId))
	log.Println("[DEBUG] Creating job property:", d.Id())
	return resourceJobPropertyRead(d, m, createJobProperty)
}

func resourceJobPropertyRead(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobName, propertyId, err := resourceJobPropertyId(d.Id())
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
	property := createJobProperty().fromClientJobProperty(clientProperty)

	// property := createJobProperty().(jobProperty)
	log.Println("[INFO] Updating from job property", d.Id())
	return property.setResourceData(d)
}

func resourceJobPropertyDelete(d *schema.ResourceData, m interface{}, createJobProperty func() jobProperty) error {
	jobName, propertyId, err := resourceJobPropertyId(d.Id())
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

	err = j.DeleteProperty(propertyId)
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
