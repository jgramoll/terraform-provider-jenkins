package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobParameterDefinitionChoice struct {
	Name        string    `mapstructure:"name"`
	Description string    `mapstructure:"description"`
	Choices     *[]string `mapstructure:"choices"`
}

func newJobParameterDefinitionChoice() *jobParameterDefinitionChoice {
	return &jobParameterDefinitionChoice{
		Choices: &[]string{},
	}
}

func (c *jobParameterDefinitionChoice) toClientJobParameterDefinition(id string) client.JobParameterDefinition {
	clientDefinition := client.NewJobParameterDefinitionChoice()
	clientDefinition.Id = id
	clientDefinition.Name = c.Name
	clientDefinition.Description = c.Description
	*clientDefinition.Choices.Items.Items = *c.Choices
	return clientDefinition
}

func (c *jobParameterDefinitionChoice) fromClientJobParameterDefintion(
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

func (c *jobParameterDefinitionChoice) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("name", c.Name); err != nil {
		return err
	}
	if err := d.Set("description", c.Description); err != nil {
		return err
	}
	return d.Set("choices", c.Choices)
}
