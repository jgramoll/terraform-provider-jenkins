package client

import (
	"encoding/xml"
)

type GitSCM struct {
	XMLName xml.Name `xml:"scm"`
	Class   string   `xml:"class,attr"`
	Plugin  string   `xml:"plugin,attr,omitempty"`

	ConfigVersion     string                `xml:"configVersion"`
	UserRemoteConfigs *GitUserRemoteConfigs `xml:"userRemoteConfigs"`
	Branches          *GitScmBranches       `xml:"branches"`

	DoGenerateSubmoduleConfigurations bool                    `xml:"doGenerateSubmoduleConfigurations"`
	SubmoduleCfg                      *GitScmSubmodulesConfig `xml:"submoduleCfg"`
	Extensions                        *GitScmExtensions       `xml:"extensions"`
}

func NewGitScm() *GitSCM {
	return &GitSCM{
		Class: "hudson.plugins.git.GitSCM",

		UserRemoteConfigs:                 NewGitUserRemoteConfigs(),
		Branches:                          NewGitScmBranches(),
		DoGenerateSubmoduleConfigurations: false,
		Extensions:                        NewGitScmExtensions(),
	}
}
