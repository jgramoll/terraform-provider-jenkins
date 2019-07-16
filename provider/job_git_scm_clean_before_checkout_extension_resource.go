package provider

import (
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmBranchMissingDefinition
// var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

func jobGitScmCleanBeforeCheckoutExtensionResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGitScmCleanBeforeCheckoutExtensionCreate,
		Read:   resourceJobGitScmCleanBeforeCheckoutExtensionRead,
		Delete: resourceJobGitScmCleanBeforeCheckoutExtensionDelete,

		Schema: map[string]*schema.Schema{
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
				ForceNew:    true,
			},
			"scm": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the scm definition",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceJobGitScmCleanBeforeCheckoutExtensionCreate(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	extension := newJobGitScmCleanBeforeCheckoutExtension()
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
	return resourceJobGitScmCleanBeforeCheckoutExtensionRead(d, m)
}

func resourceJobGitScmCleanBeforeCheckoutExtensionRead(d *schema.ResourceData, m interface{}) error {
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

	extension := newJobGitScmCleanBeforeCheckoutExtension()
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

func resourceJobGitScmCleanBeforeCheckoutExtensionDelete(d *schema.ResourceData, m interface{}) error {

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
