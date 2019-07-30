package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritProject struct {
	CompareType string `mapstructure:"compare_type"`
	Pattern     string `mapstructure:"pattern"`
}

func newJobGerritProject() *jobGerritProject {
	return &jobGerritProject{}
}

func newJobGerritProjectFromClient(clientProject *client.JobGerritTriggerProject) *jobGerritProject {
	project := newJobGerritProject()
	project.CompareType = clientProject.CompareType.String()
	project.Pattern = clientProject.Pattern
	return project
}

func (project *jobGerritProject) toClientProject(projectId string) (*client.JobGerritTriggerProject, error) {
	clientProject := client.NewJobGerritTriggerProject()
	clientProject.Id = projectId
	compareType, err := client.ParseCompareType(project.CompareType)
	if err != nil {
		return nil, err
	}
	clientProject.CompareType = compareType
	clientProject.Pattern = project.Pattern
	return clientProject, nil
}

func (project *jobGerritProject) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("compare_type", project.CompareType); err != nil {
		return err
	}
	return d.Set("pattern", project.Pattern)
}
