package main

import (
	"log"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmExtensionCodeFunc func(client.GitScmExtension) string

var jobGitScmExtensionCodeFuncs = map[string]jobGitScmExtensionCodeFunc{}

func jobGitScmExtensionsCode(extensions *client.GitScmExtensions) string {
	code := ""
	for _, item := range *extensions.Items {
		reflectType := reflect.TypeOf(item).String()
		if codeFunc, ok := jobGitScmExtensionCodeFuncs[reflectType]; ok {
			code += codeFunc(item)
		} else {
			log.Println("[WARNING] Unknown action type:", reflectType)
		}
	}
	return code
}
