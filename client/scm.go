package client

type GitSCM struct {
	ConfigVersion     string                `xml:"configVersion"`
	UserRemoteConfigs *GitUserRemoteConfigs `xml:"userRemoteConfigs"`

	Branches                          *Branches `xml:"branches"`
	DoGenerateSubmoduleConfigurations bool      `xml:"doGenerateSubmoduleConfigurations"`
	// submoduleCfg
	// extensions
}

func NewGitScm() *GitSCM {
	return &GitSCM{
		UserRemoteConfigs: NewGitUserRemoteConfigs(),
	}
}

func (scm *GitSCM) AppendUserRemoteConfig(config *GitUserRemoteConfig) *GitUserRemoteConfigs {
	return scm.UserRemoteConfigs.Append(config)
}
