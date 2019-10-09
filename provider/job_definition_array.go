package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type interfaceJobDefinition []map[string]interface{}

type jobDefinitionInit func() jobDefinition

var jobDefinitionInitFunc = map[client.JobDefinitionType]jobDefinitionInit{}

func (definitions *interfaceJobDefinition) toClientDefinition() (client.JobDefinition, error) {
	if len(*definitions) == 0 {
		return nil, nil
	}
	definitionData := (*definitions)[0]
	definitionTypeString, ok := definitionData["type"].(string)
	if !ok {
		return nil, errors.New("Failed to deserialize job definition, missing type")
	}
	definitionType, err := client.ParseJobDefinitionType(definitionTypeString)
	if err != nil {
		return nil, err
	}
	initFunc := jobDefinitionInitFunc[definitionType]
	if initFunc == nil {
		return nil, errors.New("Failed to deserialize job definition, missing init func")
	}
	definition := initFunc()
	if err := mapstructure.Decode(definitionData, &definition); err != nil {
		return nil, err
	}
	clientDefinition, err := definition.toClientDefinition()
	if err != nil {
		return nil, err
	}
	return clientDefinition, nil
}

func (*interfaceJobDefinition) fromClientDefinition(clientDefinition client.JobDefinition) (*interfaceJobDefinition, error) {
	if clientDefinition == nil {
		return nil, nil
	}

	definitionType := clientDefinition.GetType()
	initFunc := jobDefinitionInitFunc[definitionType]
	if initFunc == nil {
		return nil, errors.New("Failed to deserialize job definition, missing init func")
	}
	definition, err := initFunc().fromClientDefinition(clientDefinition)
	if err != nil {
		return nil, err
	}
	definitionData := map[string]interface{}{}
	if err := mapstructure.Decode(definition, &definitionData); err != nil {
		return nil, err
	}
	return &interfaceJobDefinition{definitionData}, nil
}
