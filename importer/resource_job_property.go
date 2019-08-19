package main

import (
	"log"
	"reflect"
	"strconv"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureJobPropertyFunc func(client.JobProperty) error
type jobPropertyCodeFunc func(string, client.JobProperty) string
type jobPropertyImportScriptFunc func(string, string, client.JobProperty) string

var ensureJobPropertyFuncs = map[string]ensureJobPropertyFunc{}
var jobPropertyCodeFuncs = map[string]jobPropertyCodeFunc{}
var jobPropertyImportScriptFuncs = map[string]jobPropertyImportScriptFunc{}

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

func jobPropertiesCode(properties *client.JobProperties) string {
	code := ""

	for i, item := range *properties.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobPropertyCodeFuncs[reflectType]; ok {
			code += codeFunc(strconv.Itoa(i+1), item)
		} else {
			log.Println("[WARNING] Unknown property type:", reflectType)
		}
	}

	return code
}

func jobPropertiesImportScript(jobName string, properties *client.JobProperties) string {
	code := ""

	for i, item := range *properties.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobPropertyImportScriptFuncs[reflectType]; ok {
			code += codeFunc(strconv.Itoa(i+1), jobName, item)
		} else {
			log.Println("[WARNING] Unknown property type:", reflectType)
		}
	}

	return code
}
