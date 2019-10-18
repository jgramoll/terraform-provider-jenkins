package provider

import (
	"fmt"
	"reflect"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobGitScmExtensionInitFunc[client.CleanBeforeCheckoutType] = func() jobGitScmExtension {
		return newJobGitScmCleanBeforeCheckoutExtension()
	}
}

type jobGitScmCleanBeforeCheckoutExtension struct {
	Type client.GitScmExtensionType `mapstructure:"type"`
}

func newJobGitScmCleanBeforeCheckoutExtension() *jobGitScmCleanBeforeCheckoutExtension {
	return &jobGitScmCleanBeforeCheckoutExtension{
		Type: client.CleanBeforeCheckoutType,
	}
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) fromClientExtension(clientExtensionInterface client.GitScmExtension) (jobGitScmExtension, error) {
	_, ok := clientExtensionInterface.(*client.GitScmCleanBeforeCheckoutExtension)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.GitScmCleanBeforeCheckoutExtension, actually %s",
			reflect.TypeOf(clientExtensionInterface).String())
	}
	return newJobGitScmCleanBeforeCheckoutExtension(), nil
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) toClientExtension() (client.GitScmExtension, error) {
	clientExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	return clientExtension, nil
}
