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

func (scm *GitSCM) GetBranch(branchId string) (*GitScmBranchSpec, error) {
	if scm.Branches == nil || scm.Branches.Items == nil {
		return nil, ErrGitScmBranchNotFound
	}
	for _, branch := range *scm.Branches.Items {
		if branch.Id == branchId {
			return branch, nil
		}
	}
	return nil, ErrGitScmBranchNotFound
}

func (scm *GitSCM) UpdateBranch(branch *GitScmBranchSpec) error {
	for _, branchRef := range *scm.Branches.Items {
		if branchRef.Id == branch.Id {
			*branchRef = *branch
			return nil
		}
	}
	return ErrGitScmBranchNotFound
}

func (scm *GitSCM) DeleteBranch(branchId string) error {
	branches := *scm.Branches.Items
	for i, branch := range branches {
		if branch.Id == branchId {
			*scm.Branches.Items = append(branches[:i], branches[i+1:]...)
			return nil
		}
	}
	return ErrGitScmBranchNotFound
}

func (scm *GitSCM) GetUserRemoteConfig(configId string) (*GitUserRemoteConfig, error) {
	if scm.UserRemoteConfigs == nil || scm.UserRemoteConfigs.Items == nil {
		return nil, ErrGitScmUserRemoteConfigNotFound
	}
	for _, config := range *scm.UserRemoteConfigs.Items {
		if config.Id == configId {
			return config, nil
		}
	}
	return nil, ErrGitScmUserRemoteConfigNotFound
}

func (scm *GitSCM) UpdateUserRemoteConfig(config *GitUserRemoteConfig) error {
	configs := *scm.UserRemoteConfigs.Items
	for _, c := range configs {
		if c.Id == config.Id {
			*c = *config
			return nil
		}
	}
	return ErrGitScmUserRemoteConfigNotFound
}

func (scm *GitSCM) DeleteUserRemoteConfig(configId string) error {
	configs := *scm.UserRemoteConfigs.Items
	for i, c := range configs {
		if c.Id == configId {
			*scm.UserRemoteConfigs.Items = append(configs[:i], configs[i+1:]...)
			return nil
		}
	}
	return ErrGitScmUserRemoteConfigNotFound
}
