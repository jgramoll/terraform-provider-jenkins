package main

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureJobPropertyFunc func(client.JobProperty) error

var ensureJobPropertyFuncs = map[string]ensureJobPropertyFunc{}

func ensureJobProperties(properties *client.JobProperties) error {
	if properties == nil || properties.Items == nil {
		return nil
	}
	for _, item := range *properties.Items {
		if item.GetId() == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.SetId(id.String())
		}
		reflectType := reflect.TypeOf(item).String()
		if ensureFunc, ok := ensureJobPropertyFuncs[reflectType]; ok {
			if err := ensureFunc(item); err != nil {
				return err
			}
		}
	}
	return nil
}
