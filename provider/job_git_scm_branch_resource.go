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

// ErrGitScmBranchMissingDefinition
var ErrGitScmBranchMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_branch")

// ErrInvalidJobGitScmBranchId
var ErrInvalidJobGitScmBranchId = errors.New("Invalid git scm id, must be jobName_scmId_branchId")

func jobGitScmBranchResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGitScmBranchCreate,
		Read:   resourceJobGitScmBranchRead,
		Update: resourceJobGitScmBranchUpdate,
		Delete: resourceJobGitScmBranchDelete,

		Schema: map[string]*schema.Schema{
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

func resourceJobGitScmBranchId(input string) (jobName string, scmId string, branchId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidJobGitScmBranchId
		return
	}
	jobName = parts[0]
	scmId = parts[1]
	branchId = parts[2]
	return
}

func resourceJobGitScmBranchCreate(d *schema.ResourceData, m interface{}) error {
	jobName, scmId, err := resourceJobDefinitionId(d.Get("scm").(string))
	if err != nil {
		return err
	}

	branch := newJobGitScmBranch()
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

	if j.Definition == nil {
		jobLock.Unlock(jobName)
		return ErrGitScmBranchMissingDefinition
	}

	// TODO better place for this cast?
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.Branches = definition.SCM.Branches.Append(branch.toClientBranch(branchId))
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, scmId, branchId}, IdDelimiter))
	log.Println("[DEBUG] Creating job git scm branch:", d.Id())
	return resourceJobGitScmBranchRead(d, m)
}

func resourceJobGitScmBranchUpdate(d *schema.ResourceData, m interface{}) error {
	jobName, _, branchId, err := resourceJobGitScmBranchId(d.Id())

	branch := newJobGitScmBranch()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &branch); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.UpdateBranch(branch.toClientBranch(branchId))
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job git scm branch:", d.Id())
	return resourceJobGitScmBranchRead(d, m)
}

func resourceJobGitScmBranchRead(d *schema.ResourceData, m interface{}) error {
	jobName, _, branchId, err := resourceJobGitScmBranchId(d.Id())
	if err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.RLock(jobName)
	j, err := jobService.GetJob(jobName)
	jobLock.RUnlock(jobName)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	if j.Definition == nil {
		return ErrGitScmBranchMissingDefinition
	}
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	clientBranch, err := definition.SCM.GetBranch(branchId)
	if err != nil {
		log.Println("[WARN] No Gerrit Branch found:", err)
		d.SetId("")
		return nil
	}
	branch := newGitScmBranchFromClient(clientBranch)

	log.Println("[INFO] Updating state for job git scm branch", d.Id())
	return branch.setResourceData(d)
}

func resourceJobGitScmBranchDelete(d *schema.ResourceData, m interface{}) error {
	jobName, _, branchId, err := resourceJobGitScmBranchId(d.Id())

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	if j.Definition == nil {
		return ErrGitScmBranchMissingDefinition
	}
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.DeleteBranch(branchId)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
