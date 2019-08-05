package main

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobActionCodeFunc func(client.JobAction) string

var jobActionCodeFuncs = map[string]jobActionCodeFunc{}

func ensureJobActions(actions *client.JobActions) error {
	if actions == nil || actions.Items == nil {
		return nil
	}
	for _, item := range *actions.Items {
		if item.GetId() == "" {
			id, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			item.SetId(id.String())
		}
	}
	return nil
}

func jobActionsCode(actions *client.JobActions) string {
	code := ""
	for _, item := range *actions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobActionCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown action type:", reflectType)
		}
	}
	return code
}
