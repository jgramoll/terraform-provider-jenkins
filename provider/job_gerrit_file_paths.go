package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGerritFilePaths []*jobGerritFilePath

func (filePaths *jobGerritFilePaths) toClientFilePaths() (*client.JobGerritTriggerFilePaths, error) {
	clientFilePaths := client.NewJobGerritTriggerFilePaths()
	for _, filePath := range *filePaths {
		clientFilePath, err := filePath.toClientFilePath()
		if err != nil {
			return nil, err
		}
		clientFilePaths = clientFilePaths.Append(clientFilePath)
	}
	return clientFilePaths, nil
}

func (*jobGerritFilePaths) fromClientFilePaths(clientFilePaths *client.JobGerritTriggerFilePaths) *jobGerritFilePaths {
	if clientFilePaths == nil || clientFilePaths.Items == nil {
		return nil
	}
	filePaths := jobGerritFilePaths{}
	for _, clientFilePath := range *clientFilePaths.Items {
		filePath := newJobGerritFilePathFromClient(clientFilePath)
		filePaths = append(filePaths, filePath)
	}
	return &filePaths
}
