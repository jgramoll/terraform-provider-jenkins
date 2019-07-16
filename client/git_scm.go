package client

import "encoding/xml"

type GitSCM struct {
	XMLName           xml.Name              `xml:"scm"`
	Class             string                `xml:"class,attr"`
	ConfigVersion     string                `xml:"configVersion"`
	UserRemoteConfigs *GitUserRemoteConfigs `xml:"userRemoteConfigs"`
	Branches          *GitScmBranches       `xml:"branches"`

	DoGenerateSubmoduleConfigurations bool `xml:"doGenerateSubmoduleConfigurations"`
	// submoduleCfg
	// extensions
}

func NewGitScm() *GitSCM {
	return &GitSCM{
		Class:             "hudson.plugins.git.GitSCM",
		Branches:          NewGitScmBranches(),
		UserRemoteConfigs: NewGitUserRemoteConfigs(),
	}
}

func (scm *GitSCM) AppendUserRemoteConfig(config *GitUserRemoteConfig) *GitUserRemoteConfigs {
	return scm.UserRemoteConfigs.Append(config)
}

func (scm *GitSCM) AppendBranch(branch *GitScmBranchSpec) *GitScmBranches {
	return scm.Branches.Append(branch)
}
