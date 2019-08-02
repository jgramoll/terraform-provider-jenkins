package main

import (
	"github.com/google/uuid"
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
	if job.Id == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		job.Id = id.String()
	}
	if err := ensureJobActions(job.Actions); err != nil {
		return err
	}
	if err := ensureJobProperties(job.Properties); err != nil {
		return err
	}
	if err := ensureJobDefinition(job.Definition); err != nil {
		return err
	}

	return s.jobService.UpdateJob(job)
}
