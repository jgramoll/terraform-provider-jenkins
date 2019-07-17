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
	var newProjectItems []*JobGerritTriggerProject
	if projects.Items != nil {
		newProjectItems = append(*projects.Items, project)
	} else {
		newProjectItems = append(newProjectItems, project)
	}
	newProjects := NewJobGerritTriggerProjects()
	newProjects.Items = &newProjectItems
	return newProjects
}
