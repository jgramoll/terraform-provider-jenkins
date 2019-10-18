package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type interfaceJobParameters []map[string]interface{}

type jobParameterInit func() jobParameterDefinition

var jobParameterInitFunc = map[client.JobParameterDefinitionType]jobParameterInit{}

func (parameters *interfaceJobParameters) toClientParameters() (*client.JobParameterDefinitions, error) {
	clientParameters := client.NewJobParameterDefinitions()
	for _, mapData := range *parameters {
		parameterTypeString, ok := mapData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize job parameter, missing type")
		}
		parameterType, err := client.ParseJobParameterDefinitionType(parameterTypeString)
		if err != nil {
			return nil, err
		}
		initFunc := jobParameterInitFunc[parameterType]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job parameter, missing init func")
		}
		parameter := initFunc()
		if err := mapstructure.Decode(mapData, &parameter); err != nil {
			return nil, err
		}
		clientParameter, err := parameter.toClientJobParameterDefinition()
		if err != nil {
			return nil, err
		}
		clientParameters = clientParameters.Append(clientParameter)
	}
	return clientParameters, nil
}

func (*interfaceJobParameters) fromClientParameters(clientParameters *client.JobParameterDefinitions) (*interfaceJobParameters, error) {
	parameters := interfaceJobParameters{}
	if clientParameters != nil && clientParameters.Items != nil {
		for _, clientParameter := range *clientParameters.Items {
			parameterType := clientParameter.GetType()
			initFunc := jobParameterInitFunc[parameterType]
			if initFunc == nil {
				return nil, errors.New("Failed to deserialize job parameter, missing init func")
			}
			parameter, err := initFunc().fromClientJobParameterDefinition(clientParameter)
			if err != nil {
				return nil, err
			}
			mapData := map[string]interface{}{}
			if err := mapstructure.Decode(parameter, &mapData); err != nil {
				return nil, err
			}
			parameters = append(parameters, mapData)
		}
	}
	return &parameters, nil
}
