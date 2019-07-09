package client

import (
	"fmt"
	"strings"
)

// JobService for interacting with jenkins jobs
type JobService struct {
	*Client
}

type JobsResponse struct {
	Jobs *[]*Job `xml:"job"`
}

func jobNameToUrl(jobName string) string {
	nameParts := strings.Split(jobName, "/")
	return "job/" + strings.Join(nameParts[:], "/job/")
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

func (service *JobService) GetJob(jobFullName string) (*Job, error) {
	details, err := service.getJobDetails(jobFullName)
	if err != nil {
		return nil, err
	}
	config, configErr := service.getJobConfig(jobFullName)
	if configErr != nil {
		return nil, configErr
	}
	return newJobFromConfigAndDetails(config, details), nil
}

func (service *JobService) getJobDetails(jobFullName string) (*jobDetails, error) {
	path := fmt.Sprintf("/%s/api/xml", jobNameToUrl(jobFullName))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response jobDetails
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	return &response, nil
}

func (service *JobService) getJobConfig(jobFullName string) (*jobConfig, error) {
	path := fmt.Sprintf("/%s/config.xml", jobNameToUrl(jobFullName))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response jobConfig
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	return &response, nil
}

func (service *JobService) CreateJob(job *Job) error {
	path := fmt.Sprintf("/%s/createItem?name=%s", jobNameToUrl(job.Folder()), job.NameOnly())
	req, err := service.NewRequestWithBody("POST", path, JobConfigFromJob(job))
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	if respErr != nil {
		return respErr
	}

	return nil
}

func (service *JobService) UpdateJob(job *Job) error {
	path := fmt.Sprintf("/%s/config.xml", jobNameToUrl(job.Name))
	req, err := service.NewRequestWithBody("POST", path, JobConfigFromJob(job))
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	if respErr != nil {
		return respErr
	}

	return nil
}

func (service *JobService) DeleteJob(jobFullName string) error {
	path := fmt.Sprintf("/%s/doDelete", jobNameToUrl(jobFullName))
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
