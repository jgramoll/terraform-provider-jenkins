package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmUserRemoteConfigs []*jobGitScmUserRemoteConfig

func (configs *jobGitScmUserRemoteConfigs) toClientConfigs() *client.GitUserRemoteConfigs {
	clientConfigs := client.NewGitUserRemoteConfigs()
	for _, config := range *configs {
		clientConfigs = clientConfigs.Append(config.toClientConfig())
	}
	return clientConfigs
}

func (*jobGitScmUserRemoteConfigs) fromClientConfigs(clientConfigs *client.GitUserRemoteConfigs) *jobGitScmUserRemoteConfigs {
	if clientConfigs == nil || clientConfigs.Items == nil {
		return nil
	}
	configs := jobGitScmUserRemoteConfigs{}
	for _, clientConfig := range *clientConfigs.Items {
		config := newJobGitScmUserRemoteConfigFromClient(clientConfig)
		configs = append(configs, config)
	}
	return &configs
}
