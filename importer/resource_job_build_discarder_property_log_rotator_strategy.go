package main

import (
	"fmt"

	"github.com/jgramoll/terraform-provider-jenkins/client"
	"github.com/jgramoll/terraform-provider-jenkins/provider"
)

func init() {
	jobBuildDiscarderPropertyStrategyCodeFuncs["*client.JobBuildDiscarderPropertyStrategyLogRotator"] = jobBuildDiscarderPropertyLogRotatorStrategyCode
	jobBuildDiscarderPropertyStrategyImportScriptFuncs["*client.JobBuildDiscarderPropertyStrategyLogRotator"] = jobBuildDiscarderPropertyLogRotatorStrategyImportScript
}

func jobBuildDiscarderPropertyLogRotatorStrategyCode(propertyIndex int, strategyInterface client.JobBuildDiscarderPropertyStrategy) string {
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

func jobBuildDiscarderPropertyLogRotatorStrategyImportScript(
	jobName string, propertyId string, strategyInterface client.JobBuildDiscarderPropertyStrategy,
) string {
	strategy := strategyInterface.(*client.JobBuildDiscarderPropertyStrategyLogRotator)
	return fmt.Sprintf(`
terraform import jenkins_job_build_discarder_property_log_rotator_strategy.main "%v"
`, provider.ResourceJobDiscarderPropertyStrategyId(jobName, propertyId, strategy.Id))
}
