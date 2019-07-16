package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmBranchMissingDefinition
// var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

func jobGerritTriggerDraftPublishedEventResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritTriggerDraftPublishedEventCreate,
		Read:   resourceJobGerritTriggerDraftPublishedEventRead,
		Delete: resourceJobGerritTriggerDraftPublishedEventDelete,

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
			"trigger": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the trigger",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceJobGerritTriggerDraftPublishedEventCreate(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	extension := newJobGerritTriggerDraftPublishedEvent()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &extension); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// TODO
	// extension.RefId = id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	// TODO better place for this cast?
	// definition := j.Definition.(*client.CpsScmFlowDefinition)
	// definition.SCM.AppendBranch(branch.toClientExtension())
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job git scm branch:", id)
	d.SetId(id.String())
	return resourceJobGerritTriggerDraftPublishedEventRead(d, m)
}

func resourceJobGerritTriggerDraftPublishedEventRead(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.RLock(jobName)
	_, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	extension := newJobGerritTriggerDraftPublishedEvent()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &extension); err != nil {
		return err
	}

	// definition := j.Definition.(*client.CpsScmFlowDefinition)
	// if definition == nil {
	// 	return nil
	// }

	log.Println("[INFO] Updating from job git scm clean before checkout extension", extension)
	return extension.setResourceData(d)
}

func resourceJobGerritTriggerDraftPublishedEventDelete(d *schema.ResourceData, m interface{}) error {

	jobName := d.Get("job").(string)
	jobLock.Lock(jobName)

	jobService := m.(*Services).JobService
	_, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	// definition := j.Definition.(*client.CpsScmFlowDefinition)
	// if definition == nil {
	// 	return nil
	// }
	jobLock.Unlock(jobName)

	d.SetId("")
	return nil
}
