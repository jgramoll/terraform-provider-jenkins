package main

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type JobImportService struct {
	jobService *client.JobService
}

func NewJobImportService(jenkinsClient *client.Client) *JobImportService {
	return &JobImportService{
		jobService: &client.JobService{Client: jenkinsClient},
	}
}

func (s *JobImportService) Import(jobName string, skipEnsure bool, outputDir string) error {
	job, err := s.jobService.GetJob(jobName)
	if err != nil {
		return err
	}
	if !skipEnsure {
		if err := NewEnsureJobResourceService(s.jobService).EnsureResourceIds(job); err != nil {
			return err
		}
	}

	if err := NewGenerateTerraformCodeService(s.jobService).GenerateCode(job, outputDir); err != nil {
		return err
	}
	// GenerateTerraformImportScriptService(s.jobService).Generate(job)
	return nil
}
