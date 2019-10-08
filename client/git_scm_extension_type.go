package client

import (
	"errors"
	"fmt"
)

type GitScmExtensionType string

var CleanBeforeCheckoutType GitScmExtensionType = "CleanBeforeCheckout"

func ParseGitScmExtensionType(s string) (GitScmExtensionType, error) {
	switch s {
	default:
		return "", errors.New(fmt.Sprintf("Unknown Git SCM Extension Type %s", s))
	case string(CleanBeforeCheckoutType):
		return CleanBeforeCheckoutType, nil
	}
}
