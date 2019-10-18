package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type jobBuildDiscarderPropertyStrategies []map[string]interface{}

type jobBuildDiscarderPropertyStrategyInit func() jobBuildDiscarderPropertyStrategy

var jobBuildDiscarderPropertyStrategyInitFunc = map[client.JobBuildDiscarderPropertyStrategyType]jobBuildDiscarderPropertyStrategyInit{}

func (strategies *jobBuildDiscarderPropertyStrategies) toClientStrategies() (*client.JobBuildDiscarderPropertyStrategyXml, error) {
	for _, mapData := range *strategies {
		strategyTypeString, ok := mapData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize job build discarder property strategy, missing type")
		}
		strategyType, err := client.ParseJobBuildDiscarderPropertyStrategyType(strategyTypeString)
		if err != nil {
			return nil, err
		}
		initFunc := jobBuildDiscarderPropertyStrategyInitFunc[strategyType]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job build discarder property strategy, missing init func")
		}
		strategy := initFunc()
		if err := mapstructure.Decode(mapData, &strategy); err != nil {
			return nil, err
		}
		return &client.JobBuildDiscarderPropertyStrategyXml{
			Item: strategy.toClientStrategy(),
		}, nil
	}
	return nil, nil
}

func (*jobBuildDiscarderPropertyStrategies) fromClientStrategy(clientStrategyInterface *client.JobBuildDiscarderPropertyStrategyXml) (*jobBuildDiscarderPropertyStrategies, error) {
	if clientStrategyInterface == nil || clientStrategyInterface.Item == nil {
		return nil, nil
	}
	clientStrategy := clientStrategyInterface.Item
	strategyType := clientStrategy.GetType()
	initFunc := jobBuildDiscarderPropertyStrategyInitFunc[strategyType]
	if initFunc == nil {
		return nil, errors.New("Failed to deserialize job build discarder property strategy, missing init func")
	}
	strategy, err := initFunc().fromClientStrategy(clientStrategy)
	if err != nil {
		return nil, err
	}
	mapData := map[string]interface{}{}
	if err := mapstructure.Decode(strategy, &mapData); err != nil {
		return nil, err
	}
	return &jobBuildDiscarderPropertyStrategies{mapData}, nil
}
