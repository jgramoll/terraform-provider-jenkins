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

// ErrInvalidTriggerGerritProjectId
var ErrInvalidTriggerGerritProjectId = errors.New("Invalid trigger gerrit project id, must be jobName_propertyId_triggerId_projectId")

func jobGerritProjectResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritProjectCreate,
		Update: resourceJobGerritProjectUpdate,
		Read:   resourceJobGerritProjectRead,
		Delete: resourceJobGerritProjectDelete,

		Schema: map[string]*schema.Schema{
			"trigger": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the Trigger",
				Required:    true,
				ForceNew:    true,
			},
			"compare_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of strategy to use for matching gerrit project",
				Required:    true,
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Pattern to use for matching gerrit project",
				Required:    true,
			},
		},
	}
}

func resourceJobGerritProjectId(input string) (jobName string, propertyId string, triggerId string, projectId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 4 {
		err = ErrInvalidTriggerGerritProjectId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	triggerId = parts[2]
	projectId = parts[3]
	return
}

func resourceJobGerritProjectCreate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, err := resourceJobTriggerId(d.Get("trigger").(string))

	project := newJobGerritProject()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &project); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	projectId := id.String()

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
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	clientProject, err := project.toClientProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	trigger.Projects = trigger.Projects.Append(clientProject)

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, propertyId, triggerId, projectId}, IdDelimiter))
	log.Println("[DEBUG] Creating job trigger gerrit project:", d.Id())
	return resourceJobGerritProjectRead(d, m)
}

func resourceJobGerritProjectUpdate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, err := resourceJobGerritProjectId(d.Id())

	project := newJobGerritProject()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &project); err != nil {
		return err
	}
	clientProject, err := project.toClientProject(projectId)
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
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	err = trigger.UpdateProject(clientProject)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	return resourceJobGerritProjectRead(d, m)
}

func resourceJobGerritProjectRead(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, err := resourceJobGerritProjectId(d.Id())

	jobService := m.(*Services).JobService
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
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		return err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	clientProject, err := trigger.GetProject(projectId)
	if err != nil {
		return err
	}
	project := newJobGerritProjectFromClient(clientProject)

	log.Println("[INFO] Updating gerrit project state from client", d.Id())
	return project.setResourceData(d)
}

func resourceJobGerritProjectDelete(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, err := resourceJobGerritProjectId(d.Id())

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
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	err = trigger.DeleteProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)

	d.SetId("")
	return nil
}
