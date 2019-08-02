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

// ErrInvalidPropertyId
var ErrInvalidPropertyStrategyId = errors.New("Invalid property strategy id, must be jobName_propertyId_strategyId")

func resourceJobPropertyStrategyId(input string) (jobName string, propertyId string, strategyId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidPropertyStrategyId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	strategyId = parts[2]
	return
}

func resourceJobPropertyStrategyImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, _, err := resourceJobPropertyStrategyId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("property", strings.Join([]string{jobName, propertyId}, IdDelimiter))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobPropertyStrategyCreate(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	jobName, propertyId, err := resourceJobPropertyId(d.Get("property").(string))
	if err != nil {
		return err
	}

	strategy := createJobBuildDiscarderPropertyStrategy()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &strategy); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	strategyId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	property, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	discardProperty := property.(*client.JobBuildDiscarderProperty)
	discardProperty.Strategy.Item = strategy.toClientStrategy(strategyId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, propertyId, strategyId}, IdDelimiter))
	log.Println("[DEBUG] Creating build discarder propety strategy", d.Id())
	return resourceJobPropertyStrategyRead(d, m, createJobBuildDiscarderPropertyStrategy)
}

func resourceJobPropertyStrategyRead(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	jobService := m.(*Services).JobService
	jobName, propertyId, err := resourceJobPropertyId(d.Get("property").(string))
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

	property, err := j.GetProperty(propertyId)
	if err != nil {
		return err
	}
	discardProperty := property.(*client.JobBuildDiscarderProperty)
	if discardProperty.Strategy.Item == nil {
		log.Println("[WARN] No Build Discarder Property Strategy found:", err)
		d.SetId("")
		return nil
	}
	strategy, err := createJobBuildDiscarderPropertyStrategy().fromClientStrategy(discardProperty.Strategy.Item)
	if err != nil {
		return err
	}
	log.Println("[INFO] Reading build discarder propety strategy", d.Id())
	return strategy.setResourceData(d)
}

func resourceJobPropertyStrategyUpdate(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	jobName, propertyId, strategyId, err := resourceJobPropertyStrategyId(d.Id())
	if err != nil {
		return err
	}

	strategy := createJobBuildDiscarderPropertyStrategy()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &strategy); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	property, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	discardProperty := property.(*client.JobBuildDiscarderProperty)

	discardProperty.Strategy.Item = strategy.toClientStrategy(strategyId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updating build discarder propety strategy", d.Id())
	return resourceJobPropertyStrategyRead(d, m, createJobBuildDiscarderPropertyStrategy)
}

func resourceJobPropertyStrategyDelete(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	jobName, propertyId, err := resourceJobPropertyId(d.Get("property").(string))
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

	property, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	discardProperty := property.(*client.JobBuildDiscarderProperty)
	discardProperty.Strategy.Item = nil
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
