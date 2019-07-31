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

// ErrInvalidJobGitScmExtensionId
var ErrInvalidJobGitScmExtensionId = errors.New("Invalid git scm extension id, must be jobName_definitionId_extensionId")

func resourceJobGitScmExtensionId(input string) (jobName string, definitionId string, extensionId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 3 {
		err = ErrInvalidJobGitScmExtensionId
		return
	}
	jobName = parts[0]
	definitionId = parts[1]
	extensionId = parts[2]
	return
}

func resourceJobGitScmExtensionCreate(d *schema.ResourceData, m interface{}, createGitScmExtension func() jobGitScmExtension) error {
	jobName, definitionId, err := resourceJobDefinitionId(d.Get("scm").(string))
	if err != nil {
		return err
	}

	extension := createGitScmExtension()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &extension); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	extensionId := id.String()

	clientExtension, err := extension.toClientExtension(extensionId)
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

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	definition.SCM.Extensions = definition.SCM.Extensions.Append(clientExtension)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, definitionId, extensionId}, IdDelimiter))
	log.Println("[DEBUG] Creating job git scm extension:", d.Id())
	return resourceJobGitScmExtensionRead(d, m, createGitScmExtension)
}

func resourceJobGitScmExtensionRead(d *schema.ResourceData, m interface{}, createGitScmExtension func() jobGitScmExtension) error {
	jobName, _, extensionId, err := resourceJobGitScmExtensionId(d.Id())
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
	definition := j.Definition.(*client.CpsScmFlowDefinition)
	clientExtension, err := definition.SCM.GetExtension(extensionId)
	if err != nil {
		return err
	}
	extension := createGitScmExtension().fromClientExtension(clientExtension)

	log.Println("[INFO] Updating state for job git scm extension", d.Id())
	return extension.setResourceData(d)
}

func resourceJobGitScmExtensionUpdate(d *schema.ResourceData, m interface{}, createGitScmExtension func() jobGitScmExtension) error {
	jobName, _, extensionId, err := resourceJobGitScmExtensionId(d.Id())
	if err != nil {
		return err
	}

	extension := createGitScmExtension()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &extension); err != nil {
		return err
	}
	clientExtension, err := extension.toClientExtension(extensionId)
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

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	err = definition.SCM.UpdateExtension(clientExtension)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updating job git scm extension", d.Id())
	return resourceJobGitScmExtensionRead(d, m, createGitScmExtension)
}

func resourceJobGitScmExtensionDelete(d *schema.ResourceData, m interface{}, createGitScmExtension func() jobGitScmExtension) error {
	jobName, _, extensionId, err := resourceJobGitScmExtensionId(d.Id())
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

	definition := j.Definition.(*client.CpsScmFlowDefinition)
	err = definition.SCM.DeleteExtension(extensionId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
