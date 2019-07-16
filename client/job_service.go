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
	Jobs *[]*jobDetails `xml:"job"`
}

func jobNameToUrl(jobName string) string {
	nameParts := strings.Split(jobName, "/")
	return "job/" + strings.Join(nameParts[:], "/job/")
}

// GetJobs get all jobs
func (service *JobService) GetJobs(folder string) (*[]*Job, error) {
	path := fmt.Sprintf("/%s/api/xml?tree=jobs[name,fullName,description]", jobNameToUrl(folder))
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var response JobsResponse
	_, respErr := service.DoWithResponse(req, &response)
	if respErr != nil {
		return nil, respErr
	}

	var jobs []*Job
	for _, job := range *response.Jobs {
		jobs = append(jobs, newJobFromConfigAndDetails(nil, job))
	}

	return &jobs, nil
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
	path := fmt.Sprintf("/%s/createItem", jobNameToUrl(job.Folder()))
	req, err := service.NewRequestWithBody("POST", path, JobConfigFromJob(job))
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("name", job.NameOnly())
	req.URL.RawQuery = q.Encode()

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
