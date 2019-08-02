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
	if err := ensureJobConfig(); err != nil {
		return err
	}
	if err := ensureJobActions(); err != nil {
		return err
	}
	if err := ensureJobProperties(); err != nil {
		return err
	}
	if err := ensureJobDefinition(); err != nil {
		return err
	}

	return s.jobService.UpdateJob(job)
	// return nil
}

func ensureJobConfig() error {
	return nil
}
func ensureJobActions() error {
	return nil
}
func ensureJobProperties() error {
	return nil
}
func ensureJobDefinition() error {
	return nil
}
