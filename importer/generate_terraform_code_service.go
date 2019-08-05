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
	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}

	if err := s.generateProviderCode(outputDir); err != nil {
		return err
	}

	if err := s.generatePipelineCode(outputDir, job); err != nil {
		return err
	}

	return nil
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

func ensureOutputDir(outputDir string) error {
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
