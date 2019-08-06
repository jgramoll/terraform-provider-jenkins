package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritFilePath struct {
	CompareType string `mapstructure:"compare_type"`
	Pattern     string `mapstructure:"pattern"`
}

func newJobGerritFilePath() *jobGerritFilePath {
	return &jobGerritFilePath{}
}

func newJobGerritFilePathFromClient(
	clientFilePath *client.JobGerritTriggerFilePath,
) *jobGerritFilePath {
	filePath := newJobGerritFilePath()
	filePath.CompareType = clientFilePath.CompareType.String()
	filePath.Pattern = clientFilePath.Pattern
	return filePath
}

func (filePath *jobGerritFilePath) toClientFilePath(filePathId string) (*client.JobGerritTriggerFilePath, error) {
	clientFilePath := client.NewJobGerritTriggerFilePath()
	clientFilePath.Id = filePathId
	compareType, err := client.ParseCompareType(filePath.CompareType)
	if err != nil {
		return nil, err
	}
	clientFilePath.CompareType = compareType
	clientFilePath.Pattern = filePath.Pattern
	return clientFilePath, nil
}

func (filePath *jobGerritFilePath) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("compare_type", filePath.CompareType); err != nil {
		return err
	}
	return d.Set("pattern", filePath.Pattern)
}
