package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func init() {
	jobBuildDiscarderPropertyStrategyCodeFuncs["*client.JobBuildDiscarderPropertyStrategyLogRotator"] = jobBuildDiscarderPropertyLogRotatorStrategyCode
}

func jobBuildDiscarderPropertyLogRotatorStrategyCode(propertyIndex string, strategyInterface client.JobBuildDiscarderPropertyStrategy) string {
	strategy := strategyInterface.(*client.JobBuildDiscarderPropertyStrategyLogRotator)
	return fmt.Sprintf(`
resource "jenkins_job_build_discarder_property_log_rotator_strategy" "main" {
	property = "${jenkins_job_build_discarder_property.property_%v.id}"

	days_to_keep          = "%v"
	num_to_keep           = "%v"
	artifact_days_to_keep = "%v"
	artifact_num_to_keep  = "%v"
}
`, propertyIndex,
		strategy.DaysToKeep, strategy.NumToKeep,
		strategy.ArtifactDaysToKeep, strategy.ArtifactNumToKeep)
}
