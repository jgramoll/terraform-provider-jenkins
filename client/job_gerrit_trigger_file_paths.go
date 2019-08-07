package client

type JobGerritTriggerFilePaths struct {
	Items *[]*JobGerritTriggerFilePath `xml:"com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.FilePath"`
}

func NewJobGerritTriggerFilePaths() *JobGerritTriggerFilePaths {
	return &JobGerritTriggerFilePaths{
		Items: &[]*JobGerritTriggerFilePath{},
	}
}

func (paths *JobGerritTriggerFilePaths) Append(path *JobGerritTriggerFilePath) *JobGerritTriggerFilePaths {
	newPaths := NewJobGerritTriggerFilePaths()
	if paths.Items != nil {
		*newPaths.Items = append(*paths.Items, path)
	} else {
		*newPaths.Items = []*JobGerritTriggerFilePath{path}
	}
	return newPaths
}
