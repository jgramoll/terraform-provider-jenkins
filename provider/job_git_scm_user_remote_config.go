package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
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

func (config *jobGitScmUserRemoteConfig) toClientConfig(configId string) *client.GitUserRemoteConfig {
	clientConfig := client.NewGitUserRemoteConfig()
	clientConfig.Id = configId
	clientConfig.Refspec = config.Refspec
	clientConfig.Url = config.Url
	clientConfig.CredentialsId = config.CredentialsId
	return clientConfig
}

func (config *jobGitScmUserRemoteConfig) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("refspec", config.Refspec); err != nil {
		return err
	}
	if err := d.Set("url", config.Url); err != nil {
		return err
	}
	return d.Set("credentials_id", config.CredentialsId)
}
