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

func (service *JobService) GetJob(jobName string) (*Job, error) {
	path := fmt.Sprintf(
		"/job/%s/api/xml",
		jobName)
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

func (service *JobService) GetJobConfig(jobName string) (*JobConfig, error) {
	path := fmt.Sprintf(
		"/job/%s/config.xml",
		jobName)
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
