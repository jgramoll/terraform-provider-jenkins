package client

import (
	"fmt"
)

// JobService for interacting with jenkins jobs
type JobService struct {
	*Client
}

type JobsResponse struct {
	Jobs *[]*Job `xml:"job"`
}

func fullName(folder string, jobName string) string {
	return fmt.Sprintf("%s/job/%s", folder, jobName)
}

// GetJobs get all jobs
func (service *JobService) GetJobs() (*[]*Job, error) {
	path := "/api/xml?tree=jobs[name]"
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response JobsResponse
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	return response.Jobs, nil
}

func (service *JobService) GetJob(folder, jobName string) (*Job, error) {
	path := fmt.Sprintf("/%s/api/xml", fullName(folder, jobName))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response Job
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	return &response, nil
}

func (service *JobService) GetJobConfig(folder string, jobName string) (*JobConfig, error) {
	path := fmt.Sprintf("/%s/config.xml", fullName(folder, jobName))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response JobConfig
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	return &response, nil
}

func (service *JobService) CreateJob(folder string, jobName string, config *JobConfig) error {
	path := fmt.Sprintf("/%s/createItem?name=%s", folder, jobName)
	req, err := service.NewRequestWithBody("POST", path, config)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	if respErr != nil {
		return respErr
	}

	return nil
}

func (service *JobService) UpdateJob(folder string, jobName string, jobConfig *JobConfig) error {
	path := fmt.Sprintf("/%s/config.xml", fullName(folder, jobName))
	req, err := service.NewRequestWithBody("POST", path, jobConfig)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	if respErr != nil {
		return respErr
	}

	return nil
}

func (service *JobService) DeleteJob(folder string, jobName string) error {
	path := fmt.Sprintf("/%s/doDelete", fullName(folder, jobName))
	req, err := service.NewRequest("POST", path)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	if respErr != nil {
		return respErr
	}

	return nil
}
