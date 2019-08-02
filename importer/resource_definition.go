package main

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureDefinitionFunc func(client.JobDefinition) error

var ensureDefinitionFuncs = map[string]ensureDefinitionFunc{}

func ensureJobDefinition(definition client.JobDefinition) error {
	if definition == nil {
		return nil
	}
	if definition.GetId() == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		definition.SetId(id.String())
	}
	reflectType := reflect.TypeOf(definition).String()
	ensureFunc, ok := ensureDefinitionFuncs[reflectType]
	if !ok {
		return fmt.Errorf("Unknown Definition Type %s", reflectType)
	}
	return ensureFunc(definition)
}
