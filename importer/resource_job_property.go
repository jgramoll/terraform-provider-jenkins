package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobPropertyCodeFunc func(client.JobProperty) string

var jobPropertyCodeFuncs = map[string]jobPropertyCodeFunc{}

func jobPropertiesCode(properties *client.JobProperties) string {
	code := ""

	for _, item := range *properties.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobPropertyCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown property type:", reflectType)
		}
	}

	return code
}
