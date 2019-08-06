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

// ErrInvalidTriggerGerritFilePathId
var ErrInvalidTriggerGerritFilePathId = errors.New("Invalid gerrit file path id, must be jobName_propertyId_triggerId_projectId_filePathId")

func jobGerritFilePathResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritFilePathCreate,
		Read:   resourceJobGerritFilePathRead,
		Update: resourceJobGerritFilePathUpdate,
		Delete: resourceJobGerritFilePathDelete,
		Importer: &schema.ResourceImporter{
			State: resourceJobGerritFilePathImporter,
		},

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the Project",
				Required:    true,
				ForceNew:    true,
			},
			"compare_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of strategy to use for matching changes based on files",
				Required:    true,
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Pattern to use for matching changes based on files",
				Required:    true,
			},
		},
	}
}

func resourceJobGerritFilePathParseId(input string) (
	jobName string, propertyId string, triggerId string,
	projectId string, filePathId string, err error,
) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 5 {
		err = ErrInvalidTriggerGerritFilePathId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	triggerId = parts[2]
	projectId = parts[3]
	filePathId = parts[4]
	return
}

func ResourceJobGerritFilePathId(jobName string, propertyId string, triggerId string, projectId string, filePathId string) string {
	return joinWithIdDelimiter(jobName, propertyId, triggerId, projectId, filePathId)
}

func resourceJobGerritFilePathImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, triggerId, projectId, _, err := resourceJobGerritFilePathParseId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("project", ResourceJobGerritProjectId(jobName, propertyId, triggerId, projectId))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobGerritFilePathCreate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, err := resourceJobGerritProjectParseId(d.Get("project").(string))

	filePath := newJobGerritFilePath()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &filePath); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	filePathId := id.String()

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
	project, err := trigger.GetProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	clientFilePath, err := filePath.toClientFilePath(filePathId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	project.FilePaths = project.FilePaths.Append(clientFilePath)

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobGerritFilePathId(jobName, propertyId, triggerId, projectId, filePathId))
	log.Println("[DEBUG] Creating job gerrit filePath:", d.Id())
	return resourceJobGerritFilePathRead(d, m)
}

func resourceJobGerritFilePathUpdate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, filePathId, err := resourceJobGerritFilePathParseId(d.Id())

	filePath := newJobGerritFilePath()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &filePath); err != nil {
		return err
	}
	clientFilePath, err := filePath.toClientFilePath(filePathId)
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
	project, err := trigger.GetProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = project.UpdateFilePath(clientFilePath)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	return resourceJobGerritFilePathRead(d, m)
}

func resourceJobGerritFilePathRead(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, filePathId, err := resourceJobGerritFilePathParseId(d.Id())

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
		log.Println("[WARN] No Job Property found:", err)
		d.SetId("")
		return nil
	}
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		return err
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	project, err := trigger.GetProject(projectId)
	if err != nil {
		return err
	}
	clientFilePath, err := project.GetFilePath(filePathId)
	if err != nil {
		log.Println("[WARN] No Gerrit File Path found:", err)
		d.SetId("")
		return nil
	}
	filePath := newJobGerritFilePathFromClient(clientFilePath)

	log.Println("[INFO] Updating gerrit filePath state from client", d.Id())
	return filePath.setResourceData(d)
}

func resourceJobGerritFilePathDelete(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, filePathId, err := resourceJobGerritFilePathParseId(d.Id())

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit File Path:", err)
		d.SetId("")
		return nil
	}

	property, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit File Path:", err)
		d.SetId("")
		return nil
	}
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit File Path:", err)
		d.SetId("")
		return nil
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	project, err := trigger.GetProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit File Path:", err)
		d.SetId("")
		return nil
	}
	err = project.DeleteFilePath(filePathId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit File Path:", err)
		d.SetId("")
		return nil
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)

	d.SetId("")
	return err
}
