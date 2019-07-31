package client

import (
	"encoding/xml"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestJobConfigSerialize(t *testing.T) {
	job := NewJob()
	job.Description = "my-desc"
	job.Actions = job.Actions.Append(NewJobDeclarativeJobAction())
	job.Actions = job.Actions.Append(NewJobDeclarativeJobPropertyTrackerAction())

	definition := NewCpsScmFlowDefinition()
	definition.SCM = NewGitScm()
	definition.SCM.ConfigVersion = "my-version"

	remoteConfig := NewGitUserRemoteConfig()
	remoteConfig.Refspec = "refspec"
	remoteConfig.Url = "url.to.here"
	remoteConfig.CredentialsId = "creds"
	definition.SCM.UserRemoteConfigs = definition.SCM.UserRemoteConfigs.Append(remoteConfig)

	scmExtension := NewGitScmCleanBeforeCheckoutExtension()
	scmExtension.Id = "extension-id"
	definition.SCM.Extensions = definition.SCM.Extensions.Append(scmExtension)

	branchSpec := NewGitScmBranchSpec()
	branchSpec.Name = "branchspec"
	definition.Id = "definition-id"
	definition.SCM.Branches = definition.SCM.Branches.Append(branchSpec)
	job.Definition = definition

	gerritBranch := NewJobGerritTriggerBranch()
	gerritBranch.CompareType = CompareTypeRegExp
	gerritBranch.Pattern = "my-branch"
	gerritProject := NewJobGerritTriggerProject()
	gerritProject.CompareType = CompareTypePlain
	gerritProject.Pattern = "my-project"
	gerritProject.Branches = gerritProject.Branches.Append(gerritBranch)
	gerritTrigger := NewJobGerritTrigger()
	gerritTrigger.Projects = gerritTrigger.Projects.Append(gerritProject)
	gerritTriggerPatchsetEvent := NewJobGerritTriggerPluginPatchsetCreatedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerPatchsetEvent)
	gerritTriggerDraftEvent := NewJobGerritTriggerPluginDraftPublishedEvent()
	gerritTrigger.TriggerOnEvents = gerritTrigger.TriggerOnEvents.Append(gerritTriggerDraftEvent)
	triggerJobProperty := NewJobPipelineTriggersProperty()
	triggerJobProperty.Id = "trigger-id"
	triggerJobProperty.Triggers = triggerJobProperty.Triggers.Append(gerritTrigger)
	job.Properties = job.Properties.Append(triggerJobProperty)

	discardPropertyStrategy := NewJobBuildDiscarderPropertyStrategyLogRotator()
	discardPropertyStrategy.DaysToKeep = 1
	discardPropertyStrategy.NumToKeep = 2
	discardPropertyStrategy.ArtifactDaysToKeep = 3
	discardPropertyStrategy.ArtifactNumToKeep = 4
	discardProperty := NewJobBuildDiscarderProperty()
	discardProperty.Id = "discard-id"
	discardProperty.Strategy.Item = discardPropertyStrategy
	job.Properties = job.Properties.Append(discardProperty)

	config := JobConfigFromJob(job)
	resultBytes, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		t.Fatalf("failed to serialize xml %s", err)
	}
	result := string(resultBytes)
	expected := `<flow-definition>
	<actions>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction></org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction>
		<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
			<jobProperties></jobProperties>
			<triggers></triggers>
			<parameters></parameters>
			<options></options>
		</org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
	</actions>
	<description>my-desc</description>
	<keepDependencies>false</keepDependencies>
	<properties>
		<org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty id="trigger-id">
			<triggers>
				<com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger>
					<gerritProjects>
						<com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.GerritProject>
							<compareType>PLAIN</compareType>
							<pattern>my-project</pattern>
							<branches>
								<com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.Branch>
									<compareType>REG_EXP</compareType>
									<pattern>my-branch</pattern>
								</com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.Branch>
							</branches>
							<disableStrictForbiddenFileVerification>false</disableStrictForbiddenFileVerification>
						</com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.data.GerritProject>
					</gerritProjects>
					<skipVote>
						<onSuccessful>false</onSuccessful>
						<onFailed>false</onFailed>
						<onUnstable>false</onUnstable>
						<onNotBuilt>false</onNotBuilt>
					</skipVote>
					<silentMode>false</silentMode>
					<silentStartMode>false</silentStartMode>
					<escapeQuotes>true</escapeQuotes>
					<nameAndEmailParameterMode>PLAIN</nameAndEmailParameterMode>
					<commitMessageParameterMode>BASE64</commitMessageParameterMode>
					<changeSubjectParameterMode>PLAIN</changeSubjectParameterMode>
					<commentTextParameterMode>BASE64</commentTextParameterMode>
					<serverName>__ANY__</serverName>
					<triggerOnEvents class="linked-list">
						<com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginPatchsetCreatedEvent>
							<excludeDrafts>false</excludeDrafts>
							<excludeTrivialRebase>false</excludeTrivialRebase>
							<excludeNoCodeChange>false</excludeNoCodeChange>
							<excludePrivateState>false</excludePrivateState>
							<excludeWipState>false</excludeWipState>
						</com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginPatchsetCreatedEvent>
						<com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent></com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.events.PluginDraftPublishedEvent>
					</triggerOnEvents>
					<dynamicTriggerConfiguration>false</dynamicTriggerConfiguration>
				</com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.GerritTrigger>
			</triggers>
		</org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
		<jenkins.model.BuildDiscarderProperty id="discard-id">
			<strategy class="hudson.tasks.LogRotator">
				<daysToKeep>1</daysToKeep>
				<numToKeep>2</numToKeep>
				<artifactDaysToKeep>3</artifactDaysToKeep>
				<artifactNumToKeep>4</artifactNumToKeep>
			</strategy>
		</jenkins.model.BuildDiscarderProperty>
	</properties>
	<definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" id="definition-id">
		<scm class="hudson.plugins.git.GitSCM">
			<configVersion>my-version</configVersion>
			<userRemoteConfigs>
				<hudson.plugins.git.UserRemoteConfig>
					<refspec>refspec</refspec>
					<url>url.to.here</url>
					<credentialsId>creds</credentialsId>
				</hudson.plugins.git.UserRemoteConfig>
			</userRemoteConfigs>
			<branches>
				<hudson.plugins.git.BranchSpec id="">
					<name>branchspec</name>
				</hudson.plugins.git.BranchSpec>
			</branches>
			<doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
			<extensions>
				<hudson.plugins.git.extensions.impl.CleanBeforeCheckout id="extension-id"></hudson.plugins.git.extensions.impl.CleanBeforeCheckout>
			</extensions>
		</scm>
		<scriptPath></scriptPath>
		<lightweight>false</lightweight>
	</definition>
	<trigger></trigger>
	<disabled>false</disabled>
</flow-definition>`
	if result != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, result, true)
		t.Fatalf("job definition not expected: %s", dmp.DiffPrettyText(diffs))
	}
}
