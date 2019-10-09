package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritProjects []*jobGerritProject

func (projects *jobGerritProjects) toClientProjects() (*client.JobGerritTriggerProjects, error) {
	clientProjects := client.NewJobGerritTriggerProjects()
	for _, project := range *projects {
		clientProject, err := project.toClientProject()
		if err != nil {
			return nil, err
		}
		clientProjects = clientProjects.Append(clientProject)
	}
	return clientProjects, nil
}

func (*jobGerritProjects) fromClientProjects(clientProjects *client.JobGerritTriggerProjects) *jobGerritProjects {
	if clientProjects == nil || clientProjects.Items == nil {
		return nil
	}
	projects := jobGerritProjects{}
	for _, clientProject := range *clientProjects.Items {
		project := newJobGerritProjectFromClient(clientProject)
		projects = append(projects, project)
	}
	return &projects
}
