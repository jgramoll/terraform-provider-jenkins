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
	Extensions *GitScmExtensions `xml:"extensions"`
}

func NewGitScm() *GitSCM {
	return &GitSCM{
		Class:             "hudson.plugins.git.GitSCM",
		UserRemoteConfigs: NewGitUserRemoteConfigs(),
		Branches:          NewGitScmBranches(),

		DoGenerateSubmoduleConfigurations: false,

		Extensions: NewGitScmExtensions(),
	}
}
