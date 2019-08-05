package provider

import (
	"strings"
)

var IdDelimiter = string(0x1F)

func joinWithIdDelimiter(args ...string) string {
	return strings.Join(args, IdDelimiter)
}
