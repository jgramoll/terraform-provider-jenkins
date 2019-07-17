package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmBranchMissingDefinition
var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

func jobGitScmBranchResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGitScmBranchCreate,
		Read:   resourceJobGitScmBranchRead,
		Update: resourceJobGitScmBranchUpdate,
		Delete: resourceJobGitScmBranchDelete,

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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the git branch",
				Optional:    true,
				Default:     "*/master",
			},
		},
	}
}

func resourceJobGitScmBranchCreate(d *schema.ResourceData, m interface{}) error {
	jobName := d.Get("job").(string)

	branch := newJobGitScmBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// branch.RefId = id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	if j.Definition == nil {
		jobLock.Unlock(jobName)
		return ErrGitScmBranchMissingDefinition
	}

	// TODO better place for this cast?
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.Branches = definition.SCM.Branches.Append(branch.toClientBranch())
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job git scm branch:", id)
	d.SetId(id.String())
	return resourceJobGitScmBranchRead(d, m)
}

func resourceJobGitScmBranchUpdate(d *schema.ResourceData, m interface{}) error {
	branch := newJobGitScmBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}

	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	// j.Definition = branch.toClientBranch()
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job git scm branch:", d.Id())
	return resourceJobGitScmBranchRead(d, m)
}

func resourceJobGitScmBranchRead(d *schema.ResourceData, m interface{}) error {
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

	branch := newJobGitScmBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}

	// definition := j.Definition.(*client.CpsScmFlowDefinition)
	// if definition == nil {
	// 	return nil
	// }

	log.Println("[INFO] Updating from job git scm branch", branch)
	return branch.setResourceData(d)
}

func resourceJobGitScmBranchDelete(d *schema.ResourceData, m interface{}) error {

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
