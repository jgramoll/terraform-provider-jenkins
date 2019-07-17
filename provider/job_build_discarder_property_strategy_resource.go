package provider

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func resourceJobBuildDiscarderPropertyStrategyCreate(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
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
	discardProperty := property.(*client.JobPipelineBuildDiscarderProperty)
	discardProperty.Strategy = strategy.toClientStrategy(strategyId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s_%s", jobName, propertyId, strategyId))
	log.Println("[DEBUG] Creating", resourceName, d.Id())
	return resourceJobBuildDiscarderPropertyStrategyRead(d, m, createJobBuildDiscarderPropertyStrategy)
}

func resourceJobBuildDiscarderPropertyStrategyRead(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
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
	discardProperty := property.(*client.JobPipelineBuildDiscarderProperty)
	// logRotatorStrategy := .(*client.JobPipelineBuildDiscarderPropertyStrategyLogRotator)
	strategy := createJobBuildDiscarderPropertyStrategy().fromClientStrategy(discardProperty.Strategy)

	log.Println("[INFO] setting state for ", resourceName, d.Id())
	return strategy.setResourceData(d)
}

func resourceJobBuildDiscarderPropertyStrategyUpdate(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	// TODO
	return resourceJobBuildDiscarderPropertyStrategyRead(d, m, createJobBuildDiscarderPropertyStrategy)
}

func resourceJobBuildDiscarderPropertyStrategyDelete(d *schema.ResourceData, m interface{}, createJobBuildDiscarderPropertyStrategy func() jobBuildDiscarderPropertyStrategy) error {
	jobName, propertyId, err := resourceJobPropertyId(d.Get("property").(string))
	if err != nil {
		return err
	}
	jobLock.Lock(jobName)

	jobService := m.(*Services).JobService
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
	// TODO better place for this cast?
	discardProperty := property.(*client.JobPipelineBuildDiscarderProperty)
	discardProperty.Strategy = nil
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
