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

// ErrInvalidTriggerGerritBranchId
var ErrInvalidTriggerGerritBranchId = errors.New("Invalid gerrit branch id, must be jobName_propertyId_triggerId_projectId_branchId")

func jobGerritBranchResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritBranchCreate,
		Read:   resourceJobGerritBranchRead,
		Update: resourceJobGerritBranchUpdate,
		Delete: resourceJobGerritBranchDelete,
		Importer: &schema.ResourceImporter{
			State: resourceJobGerritBranchImporter,
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
				Description: "Type of strategy to use for matching gerrit branch",
				Required:    true,
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Pattern to use for matching gerrit branch",
				Required:    true,
			},
		},
	}
}

func resourceJobGerritBranchParseId(input string) (jobName string, propertyId string, triggerId string, projectId string, branchId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 5 {
		err = ErrInvalidTriggerGerritBranchId
		return
	}
	jobName = parts[0]
	propertyId = parts[1]
	triggerId = parts[2]
	projectId = parts[3]
	branchId = parts[4]
	return
}

func ResourceJobGerritBranchId(jobName string, propertyId string, triggerId string, projectId string, branchId string) string {
	return joinWithIdDelimiter(jobName, propertyId, triggerId, projectId, branchId)
}

func resourceJobGerritBranchImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, propertyId, triggerId, projectId, _, err := resourceJobGerritBranchParseId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("project", ResourceJobGerritProjectId(jobName, propertyId, triggerId, projectId))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobGerritBranchCreate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, err := resourceJobGerritProjectParseId(d.Get("project").(string))

	branch := newJobGerritBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	branchId := id.String()

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
	clientBranch, err := branch.toClientBranch(branchId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	project.Branches = project.Branches.Append(clientBranch)

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobGerritBranchId(jobName, propertyId, triggerId, projectId, branchId))
	log.Println("[DEBUG] Creating job gerrit branch:", d.Id())
	return resourceJobGerritBranchRead(d, m)
}

func resourceJobGerritBranchUpdate(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, branchId, err := resourceJobGerritBranchParseId(d.Id())

	branch := newJobGerritBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}
	clientBranch, err := branch.toClientBranch(branchId)
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
	err = project.UpdateBranch(clientBranch)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	return resourceJobGerritBranchRead(d, m)
}

func resourceJobGerritBranchRead(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, branchId, err := resourceJobGerritBranchParseId(d.Id())

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
		log.Println("[WARN] No Trigger found:", err)
		d.SetId("")
		return nil
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	project, err := trigger.GetProject(projectId)
	if err != nil {
		log.Println("[WARN] No Project found:", err)
		d.SetId("")
		return nil
	}
	clientBranch, err := project.GetBranch(branchId)
	if err != nil {
		log.Println("[WARN] No Gerrit Branch found:", err)
		d.SetId("")
		return nil
	}
	branch := newJobGerritBranchFromClient(clientBranch)

	log.Println("[INFO] Updating gerrit branch state from client", d.Id())
	return branch.setResourceData(d)
}

func resourceJobGerritBranchDelete(d *schema.ResourceData, m interface{}) error {
	jobName, propertyId, triggerId, projectId, branchId, err := resourceJobGerritBranchParseId(d.Id())

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit Branch:", err)
		d.SetId("")
		return nil
	}

	property, err := j.GetProperty(propertyId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit Branch:", err)
		d.SetId("")
		return nil
	}
	triggerInterface, err := property.(*client.JobPipelineTriggersProperty).GetTrigger(triggerId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit Branch:", err)
		d.SetId("")
		return nil
	}
	trigger := triggerInterface.(*client.JobGerritTrigger)
	project, err := trigger.GetProject(projectId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit Branch:", err)
		d.SetId("")
		return nil
	}
	err = project.DeleteBranch(branchId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Gerrit Branch:", err)
		d.SetId("")
		return nil
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)

	d.SetId("")
	return err
}
