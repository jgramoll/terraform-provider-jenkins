package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmUserRemoteConfig struct {
	Refspec       string `mapstructure:"refspec"`
	Url           string `mapstructure:"url"`
	CredentialsId string `mapstructure:"credentials_id"`
}

func newJobGitScmUserRemoteConfig() *jobGitScmUserRemoteConfig {
	return &jobGitScmUserRemoteConfig{}
}

func newJobGitScmUserRemoteConfigFromClient(clientConfig *client.GitUserRemoteConfig) *jobGitScmUserRemoteConfig {
	config := newJobGitScmUserRemoteConfig()
	config.Refspec = clientConfig.Refspec
	config.Url = clientConfig.Url
	config.CredentialsId = clientConfig.CredentialsId
	return config
}

func (config *jobGitScmUserRemoteConfig) toClientConfig() *client.GitUserRemoteConfig {
	clientConfig := client.NewGitUserRemoteConfig()
	clientConfig.Refspec = config.Refspec
	clientConfig.Url = config.Url
	clientConfig.CredentialsId = config.CredentialsId
	return clientConfig
}
