package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobPropertyCodeFuncs["*client.JobDatadogJobProperty"] = jobDatadogJobPropertyCode
}

func jobDatadogJobPropertyCode(propertyInterface client.JobProperty) string {
	property := propertyInterface.(*client.JobDatadogJobProperty)
	return fmt.Sprintf(`
  property {
    type = "DatadogJobProperty"
    plugin="%s"

    emit_on_checkout = %v
  }
`, property.Plugin, property.EmitOnCheckout)
}
