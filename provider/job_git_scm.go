package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScm struct {
	Job           string `mapstructure:"job"`
	ConfigVersion string `mapstructure:"config_version"`
	ScriptPath    string `mapstructure:"script_path"`
	Lightweight   bool   `mapstructure:"lightweight"`
}

func newJobGitScm() *jobGitScm {
	return &jobGitScm{
		Lightweight: false,
	}
}

func (scm *jobGitScm) fromClientJobDefintion(clientDefinition client.JobDefinition) jobDefinition {
	if clientDefinition == nil {
		return nil
	}
	clientScmDefinition := clientDefinition.(*client.CpsScmFlowDefinition)

	definition := newJobGitScm()
	definition.ConfigVersion = clientScmDefinition.SCM.ConfigVersion
	definition.ScriptPath = clientScmDefinition.ScriptPath
	definition.Lightweight = clientScmDefinition.Lightweight
	return definition
}

func (scm *jobGitScm) setResourceData(*schema.ResourceData) error {
	return nil
}

func (scm *jobGitScm) setRefID(refId string) {
	// scm.RefId = refId
}

func (scm *jobGitScm) getRefID() string {
	return ""
	// return scm.RefId
}

func (scm *jobGitScm) toClientDefinition() client.JobDefinition {
	// TODO do we need to merge data?
	definition := client.NewCpsScmFlowDefinition()
	// definition.Id = scm.RefId
	definition.SCM = scm.toClientSCM()
	definition.ScriptPath = scm.ScriptPath
	definition.Lightweight = scm.Lightweight
	return &definition
}

func (scm *jobGitScm) toClientSCM() *client.GitSCM {
	// TODO do we need to merge data?
	clientScm := client.NewGitScm()
	clientScm.ConfigVersion = scm.ConfigVersion
	return clientScm
}
