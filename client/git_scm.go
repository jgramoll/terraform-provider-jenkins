package client

import (
	"encoding/xml"
	"errors"
)

// ErrGitScmBranchNotFound git scm branch not found
var ErrGitScmBranchNotFound = errors.New("Could not find git scm branch")

// ErrGitScmExtensionNotFound git scm extension not found
var ErrGitScmExtensionNotFound = errors.New("Could not find git scm extension")

// ErrGitScmUserRemoteConfigNotFound git scm user remote config not found
var ErrGitScmUserRemoteConfigNotFound = errors.New("Could not find git scm scm user remote config")

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
