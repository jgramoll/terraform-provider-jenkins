package client

type JobGerritTriggerProjects struct {
	Items *[]*JobGerritTriggerProject `xml:",any"`
}

func NewJobGerritTriggerProjects() *JobGerritTriggerProjects {
	return &JobGerritTriggerProjects{
		Items: &[]*JobGerritTriggerProject{},
	}
}

func (projects *JobGerritTriggerProjects) Append(project *JobGerritTriggerProject) *JobGerritTriggerProjects {
	newProjects := NewJobGerritTriggerProjects()
	if projects.Items != nil {
		*newProjects.Items = append(*projects.Items, project)
	} else {
		*newProjects.Items = []*JobGerritTriggerProject{project}
	}
	return newProjects
}
