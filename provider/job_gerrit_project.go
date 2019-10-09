package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritProject struct {
	CompareType string              `mapstructure:"compare_type"`
	Pattern     string              `mapstructure:"pattern"`
	Branches    *jobGerritBranches  `mapstructure:"branch"`
	FilePaths   *jobGerritFilePaths `mapstructure:"file_path"`
}

func newJobGerritProject() *jobGerritProject {
	return &jobGerritProject{}
}

func newJobGerritProjectFromClient(clientProject *client.JobGerritTriggerProject) *jobGerritProject {
	project := newJobGerritProject()
	project.CompareType = clientProject.CompareType.String()
	project.Pattern = clientProject.Pattern
	project.Branches = project.Branches.fromClientBranches(clientProject.Branches)
	project.FilePaths = project.FilePaths.fromClientFilePaths(clientProject.FilePaths)
	return project
}

func (project *jobGerritProject) toClientProject() (*client.JobGerritTriggerProject, error) {
	clientProject := client.NewJobGerritTriggerProject()
	compareType, err := client.ParseCompareType(project.CompareType)
	if err != nil {
		return nil, err
	}
	clientProject.CompareType = compareType
	clientProject.Pattern = project.Pattern

	branches, err := project.Branches.toClientBranches()
	if err != nil {
		return nil, err
	}
	clientProject.Branches = branches

	filePaths, err := project.FilePaths.toClientFilePaths()
	if err != nil {
		return nil, err
	}
	clientProject.FilePaths = filePaths

	return clientProject, nil
}

func (project *jobGerritProject) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("compare_type", project.CompareType); err != nil {
		return err
	}
	if err := d.Set("pattern", project.CompareType); err != nil {
		return err
	}
	if err := d.Set("branch", project.Branches); err != nil {
		return err
	}
	if err := d.Set("file_path", project.FilePaths); err != nil {
		return err
	}
	return nil
}
