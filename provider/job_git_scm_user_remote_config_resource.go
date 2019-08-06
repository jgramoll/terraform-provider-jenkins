package provider

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

// ErrGitScmUserRemoteConfigMissingDefinition
var ErrGitScmUserRemoteConfigMissingDefinition = errors.New("definition must be provided for jenkins_git_scm_user_remote_config")

// ErrInvalidJobGitScmUserRemoteConfigId
var ErrInvalidJobGitScmUserRemoteConfigId = errors.New("Invalid git scm user remote config id, must be jobName_definitionId_configId")

func jobGitScmUserRemoteConfigResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobGitScmUserRemoteConfigCreate,
		Read:   resourceJobGitScmUserRemoteConfigRead,
		Update: resourceJobGitScmUserRemoteConfigUpdate,
		Delete: resourceJobGitScmUserRemoteConfigDelete,
		Importer: &schema.ResourceImporter{
			State: resourceJobGitScmUserRemoteConfigImporter,
		},

		Schema: map[string]*schema.Schema{
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

func resourceJobGitScmUserRemoteConfigParseId(input string) (jobName string, definitionId string, configId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidJobGitScmUserRemoteConfigId
		return
	}
	jobName = parts[0]
	definitionId = parts[1]
	configId = parts[2]
	return
}

func ResourceJobGitScmUserRemoteConfigId(jobName string, definitionId string, configId string) string {
	return joinWithIdDelimiter(jobName, definitionId, configId)
}

func resourceJobGitScmUserRemoteConfigImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	jobName, definitionId, _, err := resourceJobGitScmUserRemoteConfigParseId(d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("scm", ResourceJobDefinitionId(jobName, definitionId))
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceJobGitScmUserRemoteConfigCreate(d *schema.ResourceData, m interface{}) error {
	jobName, definitionId, err := resourceJobDefinitionParseId(d.Get("scm").(string))

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	configId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
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
	definition, ok := j.Definition.(*client.CpsScmFlowDefinition)
	if !ok {
		jobLock.Unlock(jobName)
		return fmt.Errorf("Failed to create job git scm user config, invalid definition %s found",
			reflect.TypeOf(j.Definition).String())
	}
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(c.toClientConfig(configId))
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(ResourceJobGitScmUserRemoteConfigId(jobName, definitionId, configId))
	log.Println("[DEBUG] Creating job git scm user remote config:", d.Id())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigUpdate(d *schema.ResourceData, m interface{}) error {
	jobName, _, configId, err := resourceJobGitScmUserRemoteConfigParseId(d.Id())

	c := newJobGitScmUserRemoteConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &c); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	if j.Definition == nil {
		jobLock.Unlock(jobName)
		return ErrGitScmUserRemoteConfigMissingDefinition
	}

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	err = definition.SCM.UpdateUserRemoteConfig(c.toClientConfig(configId))
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated job git scm user remote config:", d.Id())
	return resourceJobGitScmUserRemoteConfigRead(d, m)
}

func resourceJobGitScmUserRemoteConfigRead(d *schema.ResourceData, m interface{}) error {
	jobName, _, configId, err := resourceJobGitScmUserRemoteConfigParseId(d.Id())

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
		return ErrGitScmUserRemoteConfigMissingDefinition
	}
	definition := j.Definition.(*client.CpsScmFlowDefinition)

	clientConfig, err := definition.SCM.GetUserRemoteConfig(configId)
	if err != nil {
		log.Println("[WARN] No Config found:", err)
		d.SetId("")
		return nil
	}

	c := newJobGitScmUserRemoteConfigFromClient(clientConfig)
	log.Println("[INFO] Updating from job git scm user remote config", c)
	return c.setResourceData(d)
}

func resourceJobGitScmUserRemoteConfigDelete(d *schema.ResourceData, m interface{}) error {
	jobName, _, configId, err := resourceJobGitScmUserRemoteConfigParseId(d.Id())

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Git User Remote Config:", err)
		d.SetId("")
		return nil
	}

	if j.Definition == nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Git User Remote Config:", err)
		d.SetId("")
		return nil
	}
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	err = definition.SCM.DeleteUserRemoteConfig(configId)
	if err != nil {
		jobLock.Unlock(jobName)
		log.Println("[WARN] Could not delete Git User Remote Config:", err)
		d.SetId("")
		return nil
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
