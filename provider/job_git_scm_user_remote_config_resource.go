package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmUserRemoteConfigMissingDefinition
var ErrGitScmUserRemoteConfigMissingDefinition = errors.New("definition must be provided for git_scm_user_remote_config")

func jobGitScmUserRemoteConfigResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGitScmUserRemoteConfigCreate,
		Read:   resourceJobGitScmUserRemoteConfigRead,
		Update: resourceJobGitScmUserRemoteConfigUpdate,
		Delete: resourceJobGitScmUserRemoteConfigDelete,

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
			"refspec": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Refspec of the git commit to checkout",
				Optional:    true,
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Url to the git repo",
				Required:    true,
			},
			"credentials_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the Jenkins Credentials to use to checkout git repo",
				Optional:    true,
			},
		},
	}
}

func resourceJobGitScmUserRemoteConfigCreate(d *schema.ResourceData, m interface{}) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	c.RefId = id.String()

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	if j.Definition == nil {
		return ErrGitScmUserRemoteConfigMissingDefinition
	}

	// TODO better place for this cast?
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.AppendUserRemoteConfig(c.toClientConfig())
	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job git scm user remote config:", id)
	d.SetId(id.String())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigUpdate(d *schema.ResourceData, m interface{}) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	j.Definition = c.toClientConfig()
	err = jobService.UpdateJob(j)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job definition:", d.Id())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigRead(d *schema.ResourceData, m interface{}) error {
	jobId := d.Get("job").(string)
	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(jobId)
	if err != nil {
		log.Println("[WARN] No Job found:", err)
		d.SetId("")
		return nil
	}

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	if definition == nil {
		return nil
	}

	log.Println("[INFO] Updating from job git scm user remote config", c)
	return c.setResourceData(d)
}

func resourceJobGitScmUserRemoteConfigDelete(d *schema.ResourceData, m interface{}) error {
	jobLock.Lock()
	defer jobLock.Unlock()

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		return err
	}

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	if definition == nil {
		return nil
	}

	d.SetId("")
	return nil
}
