package client

type GitUserRemoteConfig struct {
	Refspec       string `xml:"refspec"`
	Url           string `xml:"url"`
	CredentialsId string `xml:"credentialsId"`
}

type GitUserRemoteConfigs struct {
	Items *[]*GitUserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}

func NewGitUserRemoteConfigs() *GitUserRemoteConfigs {
	return &GitUserRemoteConfigs{
		Items: &[]*GitUserRemoteConfig{},
	}
}

func (configs *GitUserRemoteConfigs) Append(config *GitUserRemoteConfig) *GitUserRemoteConfigs {
	var newConfigItems []*GitUserRemoteConfig
	if configs.Items != nil {
		newConfigItems = append(*configs.Items, config)
	} else {
		newConfigItems = append(newConfigItems, config)
	}
	newConfigs := NewGitUserRemoteConfigs()
	newConfigs.Items = &newConfigItems
	return newConfigs
}
