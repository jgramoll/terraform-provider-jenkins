package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

// ErrInvalidJobActionId
var ErrInvalidJobActionId = errors.New("Invalid action id, must be jobName\xactionId")

func resourceJobActionId(input string) (jobName string, actionId string, err error) {
	parts := strings.Split(input, IdDelimiter)
	if len(parts) != 2 {
		err = ErrInvalidJobActionId
		return
	}
	jobName = parts[0]
	actionId = parts[1]
	return
}

func resourceJobActionCreate(d *schema.ResourceData, m interface{}, createJobAction func() jobAction) error {
	jobName := d.Get("job").(string)

	action := createJobAction()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &action); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	actionId := id.String()

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	clientAction, err := action.toClientAction(actionId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	j.Actions = j.Actions.Append(clientAction)
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{jobName, actionId}, IdDelimiter))
	log.Println("[DEBUG] Creating job action:", d.Id())
	return resourceJobActionRead(d, m, createJobAction)
}

func resourceJobActionRead(d *schema.ResourceData, m interface{}, createJobAction func() jobAction) error {
	jobName, actionId, err := resourceJobActionId(d.Id())
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
	clientAction, err := j.GetAction(actionId)
	if err != nil {
		return err
	}
	action, err := createJobAction().fromClientAction(clientAction)
	if err != nil {
		return err
	}

	log.Println("[INFO] Updating state for job action", d.Id())
	return action.setResourceData(d)
}

func resourceJobActionUpdate(d *schema.ResourceData, m interface{}, createJobAction func() jobAction) error {
	jobName, actionId, err := resourceJobActionId(d.Id())
	if err != nil {
		return err
	}

	action := createJobAction()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &action); err != nil {
		return err
	}

	jobService := m.(*Services).JobService
	jobLock.Lock(jobName)
	j, err := jobService.GetJob(jobName)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}

	clientAction, err := action.toClientAction(actionId)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = j.UpdateAction(clientAction)
	if err != nil {
		jobLock.Unlock(jobName)
		return err
	}
	err = jobService.UpdateJob(j)
	jobLock.Unlock(jobName)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updating job action", d.Id())
	return resourceJobActionRead(d, m, createJobAction)
}

func resourceJobActionDelete(d *schema.ResourceData, m interface{}, createJobAction func() jobAction) error {
	jobName, actionId, err := resourceJobActionId(d.Id())
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

	err = j.DeleteAction(actionId)
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
