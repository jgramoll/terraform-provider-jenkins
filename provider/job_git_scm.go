package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScm struct {
	Plugin        string `mapstructure:"plugin"`
	GitPlugin     string `mapstructure:"git_plugin"`
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
	definition.Plugin = clientScmDefinition.Plugin
	definition.GitPlugin = clientScmDefinition.SCM.Plugin
	definition.ConfigVersion = clientScmDefinition.SCM.ConfigVersion
	definition.ScriptPath = clientScmDefinition.ScriptPath
	definition.Lightweight = clientScmDefinition.Lightweight
	return definition, nil
}

func (scm *jobGitScm) toClientDefinition(definitionId string) client.JobDefinition {
	definition := client.NewCpsScmFlowDefinition()
	definition.Id = definitionId
	definition.Plugin = scm.Plugin
	definition.SCM = scm.toClientSCM()
	definition.ScriptPath = scm.ScriptPath
	definition.Lightweight = scm.Lightweight
	return definition
}

func (scm *jobGitScm) toClientSCM() *client.GitSCM {
	clientScm := client.NewGitScm()
	clientScm.Plugin = scm.GitPlugin
	clientScm.ConfigVersion = scm.ConfigVersion
	return clientScm
}

func (scm *jobGitScm) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("plugin", scm.Plugin); err != nil {
		return err
	}
	if err := d.Set("git_plugin", scm.GitPlugin); err != nil {
		return err
	}
	if err := d.Set("config_version", scm.ConfigVersion); err != nil {
		return err
	}
	if err := d.Set("script_path", scm.ScriptPath); err != nil {
		return err
	}
	return d.Set("lightweight", scm.Lightweight)
}
