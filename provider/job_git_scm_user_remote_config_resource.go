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
var ErrGitScmUserRemoteConfigMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_user_remote_config")

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
	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// c.RefId = id.String()

	jobName := d.Get("job").(string)
	jobLock.Lock(jobName)

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	if j.Definition == nil {
		jobLock.Unlock(jobName)
		return ErrGitScmUserRemoteConfigMissingDefinition
	}

	// TODO better place for this cast?
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(c.toClientConfig())
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating job git scm user remote config:", id)
	d.SetId(id.String())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigUpdate(d *schema.ResourceData, m interface{}) error {
	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	jobName := d.Get("job").(string)
	jobLock.Lock(jobName)

	jobService := m.(*Services).JobService
	j, err := jobService.GetJob(d.Get("job").(string))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	// j.Definition = c.toClientConfig()
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job definition:", d.Id())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigRead(d *schema.ResourceData, m interface{}) error {
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

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	// definition := j.Definition.(*client.CpsScmFlowDefinition)
	// if definition == nil {
	// 	return nil
	// }

	log.Println("[INFO] Updating from job git scm user remote config", c)
	return c.setResourceData(d)
}

func resourceJobGitScmUserRemoteConfigDelete(d *schema.ResourceData, m interface{}) error {

	jobName := d.Get("job").(string)

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	_, err := jobService.GetJob(d.Get("job").(string))
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
