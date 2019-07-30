package client

type GitUserRemoteConfigs struct {
	Items *[]*GitUserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}

func NewGitUserRemoteConfigs() *GitUserRemoteConfigs {
	return &GitUserRemoteConfigs{
		Items: &[]*GitUserRemoteConfig{},
	}
}

func (configs *GitUserRemoteConfigs) Append(config *GitUserRemoteConfig) *GitUserRemoteConfigs {
	newConfigs := NewGitUserRemoteConfigs()
	if configs.Items != nil {
		*newConfigs.Items = append(*configs.Items, config)
	} else {
		*newConfigs.Items = []*GitUserRemoteConfig{config}
	}
	return newConfigs
}
