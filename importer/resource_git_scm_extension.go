package main

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmExtensionCodeFunc func(client.GitScmExtension) string

var jobGitScmExtensionCodeFuncs = map[string]jobGitScmExtensionCodeFunc{}

func ensureGitScmExtensions(extensions *client.GitScmExtensions) error {
	if extensions == nil || extensions.Items == nil {
		return nil
	}
	for _, item := range *extensions.Items {
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
