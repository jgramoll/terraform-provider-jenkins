package main

import (
	"fmt"
	"os"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type GenerateTerraformCodeService struct {
	jobService *client.JobService
}

func NewGenerateTerraformCodeService(jobService *client.JobService) *GenerateTerraformCodeService {
	return &GenerateTerraformCodeService{
		jobService: jobService,
	}
}

func (s *GenerateTerraformCodeService) GenerateCode(job *client.Job, outputDir string) error {
	if err := s.generateProviderCode(outputDir); err != nil {
		return err
	}

	return s.generatePipelineCode(outputDir, job)
}

func (s *GenerateTerraformCodeService) generateProviderCode(outputDir string) error {
	tfCodeFile, err := os.Create(fmt.Sprintf("%s/provider.tf", outputDir))
	if err != nil {
		return err
	}
	defer tfCodeFile.Close()

	_, err = tfCodeFile.Write([]byte(providerCode(s.jobService.Config.Address)))
	return err
}

func (s *GenerateTerraformCodeService) generatePipelineCode(outputDir string, job *client.Job) error {
	tfCodeFile, err := os.Create(fmt.Sprintf("%s/pipeline.tf", outputDir))
	if err != nil {
		return err
	}
	defer tfCodeFile.Close()

	_, err = tfCodeFile.Write([]byte(jobCode(job)))
	return err
}
