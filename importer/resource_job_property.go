package main

import (
	"log"
	"reflect"
	"strconv"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type ensureJobPropertyFunc func(client.JobProperty) error
type jobPropertyCodeFunc func(string, client.JobProperty) string
type jobPropertyImportScriptFunc func(string, string, client.JobProperty) string

var ensureJobPropertyFuncs = map[string]ensureJobPropertyFunc{}
var jobPropertyCodeFuncs = map[string]jobPropertyCodeFunc{}
var jobPropertyImportScriptFuncs = map[string]jobPropertyImportScriptFunc{}

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
