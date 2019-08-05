package main

import (
	"fmt"
	"os"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type GenerateTerraformImportService struct {
	jobService *client.JobService
}

func NewGenerateTerraformImportService(jobService *client.JobService) *GenerateTerraformImportService {
	return &GenerateTerraformImportService{
		jobService: jobService,
	}
}

func (s *GenerateTerraformImportService) GenerateScript(job *client.Job, outputDir string) error {
	scriptFilePath := fmt.Sprintf("%s/import.sh", outputDir)
	scriptFile, err := os.Create(scriptFilePath)
	if err != nil {
		return err
	}
	defer scriptFile.Close()

	if err := os.Chmod(scriptFilePath, 0777); err != nil {
		return err
	}

	_, err = scriptFile.Write([]byte(jobImportScript(job)))
	return err
}
