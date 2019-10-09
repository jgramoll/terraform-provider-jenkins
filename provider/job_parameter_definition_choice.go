package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobParameterInitFunc[client.ChoiceParameterDefinitionType] = func() jobParameterDefinition {
		return newJobParameterDefinitionChoice()
	}
}

type jobParameterDefinitionChoice struct {
	Type        client.JobParameterDefinitionType `mapstructure:"type"`
	Name        string                            `mapstructure:"name"`
	Description string                            `mapstructure:"description"`
	Choices     *[]string                         `mapstructure:"choices"`
}

func newJobParameterDefinitionChoice() *jobParameterDefinitionChoice {
	return &jobParameterDefinitionChoice{
		Type:    client.ChoiceParameterDefinitionType,
		Choices: &[]string{},
	}
}

func (c *jobParameterDefinitionChoice) toClientJobParameterDefinition() (client.JobParameterDefinition, error) {
	clientDefinition := client.NewJobParameterDefinitionChoice()
	clientDefinition.Name = c.Name
	clientDefinition.Description = c.Description
	*clientDefinition.Choices.Items.Items = *c.Choices
	return clientDefinition, nil
}

func (c *jobParameterDefinitionChoice) fromClientJobParameterDefinition(
	clientDefinitionInterface client.JobParameterDefinition,
) (jobParameterDefinition, error) {

	clientDefinition, ok := clientDefinitionInterface.(*client.JobParameterDefinitionChoice)
	if !ok {
		return nil, fmt.Errorf("Parameter Definition is not of expected type, expected *client.JobParameterDefinitionChoice, actually %s",
			reflect.TypeOf(clientDefinitionInterface).String())
	}
	newChoice := newJobParameterDefinitionChoice()
	newChoice.Name = clientDefinition.Name
	newChoice.Description = clientDefinition.Description
	*newChoice.Choices = *clientDefinition.Choices.Items.Items
	return newChoice, nil
}
