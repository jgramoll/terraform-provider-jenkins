package main

import (
	"fmt"
	"strings"
)

func providerCode(jenkinsAddress string) string {
	return strings.TrimSpace(fmt.Sprintf(`
provider "jenkins" {
  address = "%s"
}
`, jenkinsAddress))
}
