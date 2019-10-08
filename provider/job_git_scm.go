package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScm struct {
	Type              string                      `mapstructure:"type"`
	Plugin            string                      `mapstructure:"plugin"`
	ConfigVersion     string                      `mapstructure:"config_version"`
	Branches          *jobGitScmBranches          `mapstructure:"branch"`
	Extensions        *jobGitScmExtensions        `mapstructure:"extension"`
	UserRemoteConfigs *jobGitScmUserRemoteConfigs `mapstructure:"user_remote_config"`
}

func newJobGitScm() *jobGitScm {
	return &jobGitScm{
		Type: "GitSCM",
	}
}

func (scm *jobGitScm) toClientSCM() (*client.GitSCM, error) {
	clientScm := client.NewGitScm()
	clientScm.Plugin = scm.Plugin
	clientScm.ConfigVersion = scm.ConfigVersion
	clientScm.Branches = scm.Branches.toClientBranches()
	extensions, err := scm.Extensions.toClientExtensions()
	if err != nil {
		return nil, err
	}
	clientScm.Extensions = extensions
	clientScm.UserRemoteConfigs = scm.UserRemoteConfigs.toClientConfigs()
	return clientScm, nil
}

func (*jobGitScm) fromClientSCM(clientSCM *client.GitSCM) (*jobGitScm, error) {
	scm := newJobGitScm()
	scm.Plugin = clientSCM.Plugin
	scm.ConfigVersion = clientSCM.ConfigVersion
	scm.Branches = scm.Branches.fromClientBranches(clientSCM.Branches)
	extensions, err := scm.Extensions.fromClientExtensions(clientSCM.Extensions)
	if err != nil {
		return nil, err
	}
	scm.Extensions = extensions
	scm.UserRemoteConfigs = scm.UserRemoteConfigs.fromClientConfigs(clientSCM.UserRemoteConfigs)
	return scm, nil
}
