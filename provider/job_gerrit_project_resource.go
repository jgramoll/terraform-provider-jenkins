package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmBranchMissingDefinition
// var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

func jobGerritBranchResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGerritBranchCreate,
		Update: resourceJobGerritBranchUpdate,
		Read:   resourceJobGerritBranchRead,
		Delete: resourceJobGerritBranchDelete,

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the Trigger",
				Required:    true,
				ForceNew:    true,
			},
			"compare_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of strategy to use for discarding job history",
				Required:    true,
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of strategy to use for discarding job history",
				Required:    true,
			},
		},
	}
}

func resourceJobGerritBranchCreate(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	extension := newJobGerritProject()
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
	return resourceJobGerritBranchRead(d, m)
}

func resourceJobGerritBranchUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceJobGerritBranchRead(d, m)
}

func resourceJobGerritBranchRead(d *schema.ResourceData, m interface{}) error {
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

	extension := newJobGerritProject()
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

func resourceJobGerritBranchDelete(d *schema.ResourceData, m interface{}) error {

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
