package client

type JobGerritTriggerBranches struct {
	Items *[]*JobGerritTriggerBranch `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.Branch"`
}

func NewJobGerritTriggerBranches() *JobGerritTriggerBranches {
	return &JobGerritTriggerBranches{
		Items: &[]*JobGerritTriggerBranch{},
	}
}

func (branches *JobGerritTriggerBranches) Append(project *JobGerritTriggerBranch) *JobGerritTriggerBranches {
	var newBranchItems []*JobGerritTriggerBranch
	if branches.Items != nil {
		newBranchItems = append(*branches.Items, project)
	} else {
		newBranchItems = append(newBranchItems, project)
	}
	newBranches := NewJobGerritTriggerBranches()
	newBranches.Items = &newBranchItems
	return newBranches
}
