package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmCleanBeforeCheckoutExtension struct {
	Scm string `mapstructure:"scm"`
}

func newJobGitScmCleanBeforeCheckoutExtension() *jobGitScmCleanBeforeCheckoutExtension {
	return &jobGitScmCleanBeforeCheckoutExtension{}
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) fromClientExtension(clientExtensionInterface client.GitScmExtension) (jobGitScmExtension, error) {
	_, ok := clientExtensionInterface.(*client.GitScmCleanBeforeCheckoutExtension)
	if !ok {
		return nil, fmt.Errorf("Strategy is not of expected type, expected *client.GitScmCleanBeforeCheckoutExtension, actually %s",
			reflect.TypeOf(clientExtensionInterface).String())
	}
	return newJobGitScmCleanBeforeCheckoutExtension(), nil
}

func (branch *jobGitScmCleanBeforeCheckoutExtension) toClientExtension(extensionId string) (client.GitScmExtension, error) {
	clientExtension := client.NewGitScmCleanBeforeCheckoutExtension()
	clientExtension.Id = extensionId
	return clientExtension, nil
}

func (config *jobGitScmCleanBeforeCheckoutExtension) setResourceData(d *schema.ResourceData) error {
	return nil
}
