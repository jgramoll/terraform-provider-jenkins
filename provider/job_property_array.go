package provider

import (
	"errors"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/mitchellh/mapstructure"
)

type interfaceJobProperties []map[string]interface{}

type jobPropertyInit func() jobProperty

var jobPropertyInitFunc = map[client.JobPropertyType]jobPropertyInit{}

func (properties *interfaceJobProperties) toClientProperties() (*client.JobProperties, error) {
	clientProperties := client.NewJobProperties()
	for _, mapData := range *properties {
		propertyTypeString, ok := mapData["type"].(string)
		if !ok {
			return nil, errors.New("Failed to deserialize job property, missing type")
		}
		propertyType, err := client.ParseJobPropertyType(propertyTypeString)
		if err != nil {
			return nil, err
		}
		initFunc := jobPropertyInitFunc[propertyType]
		if initFunc == nil {
			return nil, errors.New("Failed to deserialize job property, missing init func")
		}
		property := initFunc()
		if err := mapstructure.Decode(mapData, &property); err != nil {
			return nil, err
		}
		clientProperty, err := property.toClientProperty()
		if err != nil {
			return nil, err
		}
		clientProperties = clientProperties.Append(clientProperty)
	}
	return clientProperties, nil
}

func (*interfaceJobProperties) fromClientProperties(clientProperties *client.JobProperties) (*interfaceJobProperties, error) {
	properties := interfaceJobProperties{}
	if clientProperties != nil && clientProperties.Items != nil {
		for _, clientProperty := range *clientProperties.Items {
			propertyType := clientProperty.GetType()
			initFunc := jobPropertyInitFunc[propertyType]
			if initFunc == nil {
				return nil, errors.New("Failed to deserialize job property, missing init func")
			}
			property, err := initFunc().fromClientProperty(clientProperty)
			if err != nil {
				return nil, err
			}
			mapData := map[string]interface{}{}
			if err := mapstructure.Decode(property, &mapData); err != nil {
				return nil, err
			}
			properties = append(properties, mapData)
		}
	}
	return &properties, nil
}
