package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobDefinitionInitFunc[client.CpsScmFlowDefinitionType] = func() jobDefinition {
		return newJobCpsScmFlowDefinition()
	}
}

type jobCpsScmFlowDefinition struct {
	Type        string          `mapstructure:"type"`
	Plugin      string          `mapstructure:"plugin"`
	SCM         *jobGitScmArray `mapstructure:"scm"`
	ScriptPath  string          `mapstructure:"script_path"`
	Lightweight bool            `mapstructure:"lightweight"`
}

func newJobCpsScmFlowDefinition() *jobCpsScmFlowDefinition {
	return &jobCpsScmFlowDefinition{
		Type:        string(client.CpsScmFlowDefinitionType),
		Lightweight: false,
	}
}

func (*jobCpsScmFlowDefinition) fromClientDefinition(clientDefinitionInterface client.JobDefinition) (jobDefinition, error) {
	clientDefinition, ok := clientDefinitionInterface.(*client.CpsScmFlowDefinition)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.CpsScmFlowDefinition, actually %s",
			reflect.TypeOf(clientDefinitionInterface).String())
	}

	definition := newJobCpsScmFlowDefinition()
	definition.Plugin = clientDefinition.Plugin
	scm, err := definition.SCM.fromClientSCM(clientDefinition.SCM)
	if err != nil {
		return nil, err
	}
	definition.SCM = scm
	definition.ScriptPath = clientDefinition.ScriptPath
	definition.Lightweight = clientDefinition.Lightweight
	return definition, nil
}

func (definition *jobCpsScmFlowDefinition) toClientDefinition() (client.JobDefinition, error) {
	clientDefinition := client.NewCpsScmFlowDefinition()
	clientDefinition.Plugin = definition.Plugin
	scm, err := definition.SCM.toClientSCM()
	if err != nil {
		return nil, err
	}
	clientDefinition.SCM = scm
	clientDefinition.ScriptPath = definition.ScriptPath
	clientDefinition.Lightweight = definition.Lightweight
	return clientDefinition, nil
}
