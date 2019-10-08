package main

import (
	"os"

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

func (s *JobImportService) Import(jobName string, outputDir string) error {
	job, err := s.jobService.GetJob(jobName)
	if err != nil {
		return err
	}

	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}

	if err := NewGenerateTerraformCodeService(s.jobService).GenerateCode(job, outputDir); err != nil {
		return err
	}
	return NewGenerateTerraformImportService(s.jobService).GenerateScript(job, outputDir)
}

func ensureOutputDir(outputDir string) error {
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
