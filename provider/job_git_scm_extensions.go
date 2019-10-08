package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type jobGitScmExtensions []map[string]interface{}

type jobGitScmExtensionInit func() jobGitScmExtension

var jobGitScmExtensionInitFunc = map[client.GitScmExtensionType]jobGitScmExtensionInit{}

func (extensions *jobGitScmExtensions) toClientExtensions() (*client.GitScmExtensions, error) {
	clientExtensions := client.NewGitScmExtensions()
	for _, extensionData := range *extensions {
		extensionTypeString, ok := extensionData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to parse Git SCM Extension, invalid type")
		}
		extensionType, err := client.ParseGitScmExtensionType(extensionTypeString)
		if err != nil {
			return nil, err
		}
		extension := jobGitScmExtensionInitFunc[extensionType]()
		if err := mapstructure.Decode(extensionData, &extension); err != nil {
			return nil, err
		}
		clientExtension, err := extension.toClientExtension()
		if err != nil {
			return nil, err
		}
		clientExtensions = clientExtensions.Append(clientExtension)
	}
	return clientExtensions, nil
}

func (*jobGitScmExtensions) fromClientExtensions(clientExtensions *client.GitScmExtensions) (*jobGitScmExtensions, error) {
	if clientExtensions == nil || clientExtensions.Items == nil {
		return nil, nil
	}
	extensions := jobGitScmExtensions{}
	for _, clientExtension := range *clientExtensions.Items {
		extensionType := clientExtension.GetType()
		extension := jobGitScmExtensionInitFunc[extensionType]()
		extension, err := extension.fromClientExtension(clientExtension)
		if err != nil {
			return nil, err
		}
		extensionData := map[string]interface{}{}
		if err := mapstructure.Decode(extension, &extensionData); err != nil {
			return nil, err
		}
		extensions = append(extensions, extensionData)
	}
	return &extensions, nil
}
