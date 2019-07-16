package provider

import (

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmUserRemoteConfig struct {
	RefId         string `mapstructure:"ref_id"`
	Job           string `mapstructure:"job"`
	Refspec       string `mapstructure:"refspec"`
	Url           string `mapstructure:"url"`
	CredentialsId string `mapstructure:"credentials_id"`
}

func newJobGitScmUserRemoteConfig() *jobGitScmUserRemoteConfig {
	return &jobGitScmUserRemoteConfig{}
}

func (config *jobGitScmUserRemoteConfig) toClientConfig() *client.GitUserRemoteConfig {
	return &client.GitUserRemoteConfig {
		Refspec: config.Refspec,
		Url: config.Url,
		CredentialsId: config.CredentialsId,
	}
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
