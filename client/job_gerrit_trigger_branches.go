package client

type JobGerritTriggerBranches struct {
	Items *[]*JobGerritTriggerBranch `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.Branch"`
}

func NewJobGerritTriggerBranches() *JobGerritTriggerBranches {
	return &JobGerritTriggerBranches{
		Items: &[]*JobGerritTriggerBranch{},
	}
}

func (branches *JobGerritTriggerBranches) Append(branch *JobGerritTriggerBranch) *JobGerritTriggerBranches {
	newBranches := NewJobGerritTriggerBranches()
	if branches.Items != nil {
		*newBranches.Items = append(*branches.Items, branch)
	} else {
		*newBranches.Items = []*JobGerritTriggerBranch{branch}
	}
	return newBranches
}
