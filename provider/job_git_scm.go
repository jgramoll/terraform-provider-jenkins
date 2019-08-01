package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
	"reflect"
)

type jobGitScm struct {
	ConfigVersion string `mapstructure:"config_version"`
	ScriptPath    string `mapstructure:"script_path"`
	Lightweight   bool   `mapstructure:"lightweight"`
}

func newJobGitScm() *jobGitScm {
	return &jobGitScm{
		Lightweight: false,
	}
}

func (scm *jobGitScm) fromClientJobDefintion(clientDefinitionInterface client.JobDefinition) (jobDefinition, error) {
	clientScmDefinition, ok := clientDefinitionInterface.(*client.CpsScmFlowDefinition)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.CpsScmFlowDefinition, actually %s",
			reflect.TypeOf(clientDefinitionInterface).String())
	}

	definition := newJobGitScm()
	definition.ConfigVersion = clientScmDefinition.SCM.ConfigVersion
	definition.ScriptPath = clientScmDefinition.ScriptPath
	definition.Lightweight = clientScmDefinition.Lightweight
	return definition, nil
}

func (scm *jobGitScm) setResourceData(*schema.ResourceData) error {
	return nil
}

func (scm *jobGitScm) toClientDefinition(definitionId string) client.JobDefinition {
	definition := client.NewCpsScmFlowDefinition()
	definition.Id = definitionId
	definition.SCM = scm.toClientSCM()
	definition.ScriptPath = scm.ScriptPath
	definition.Lightweight = scm.Lightweight
	return &definition
}

func (scm *jobGitScm) toClientSCM() *client.GitSCM {
	clientScm := client.NewGitScm()
	clientScm.ConfigVersion = scm.ConfigVersion
	return clientScm
}
