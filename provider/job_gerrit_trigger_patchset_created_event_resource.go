package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmBranchMissingDefinition
// var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

func jobGerritTriggerPatchSetCreatedEventResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritTriggerPatchSetCreatedEventCreate,
		Update: resourceJobGerritTriggerPatchSetCreatedEventUpdate,
		Read:   resourceJobGerritTriggerPatchSetCreatedEventRead,
		Delete: resourceJobGerritTriggerPatchSetCreatedEventDelete,

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
			"trigger": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the gerrit trigger",
				Required:    true,
			},
			"exclude_drafts": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If drafts should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_trivial_rebase": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If trivial rebase should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_no_code_change": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If no code change should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_private_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If private state should be considered for triggering job",
				Optional:    true,
				Default:     false,
			},
			"exclude_wip_state": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If wip should be excluded from triggering job",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceJobGerritTriggerPatchSetCreatedEventCreate(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	extension := newJobGerritTriggerPatchSetCreatedEvent()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &extension); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
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
	return resourceJobGerritTriggerPatchSetCreatedEventRead(d, m)
}

func resourceJobGerritTriggerPatchSetCreatedEventUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceJobGerritTriggerPatchSetCreatedEventRead(d, m)
}

func resourceJobGerritTriggerPatchSetCreatedEventRead(d *schema.ResourceData, m interface{}) error {
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

	extension := newJobGerritTriggerPatchSetCreatedEvent()
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

func resourceJobGerritTriggerPatchSetCreatedEventDelete(d *schema.ResourceData, m interface{}) error {

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
