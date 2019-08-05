package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type EnsureJobResourceService struct {
	jobService *client.JobService
}

func NewEnsureJobResourceService(jobService *client.JobService) *EnsureJobResourceService {
	return &EnsureJobResourceService{
		jobService: jobService,
	}
}

func (s *EnsureJobResourceService) EnsureResourceIds(job *client.Job) error {
	if err := ensureJob(job); err != nil {
		return err
	}

	return s.jobService.UpdateJob(job)
}
