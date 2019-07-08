package client

type SCM struct {
	ConfigVersion     int64             `xml:"configVersion"`
	UserRemoteConfigs UserRemoteConfigs `xml:"userRemoteConfigs"`

	Branches                          Branches `xml:"branches"`
	DoGenerateSubmoduleConfigurations bool     `xml:"doGenerateSubmoduleConfigurations"`
	// submoduleCfg
	//extensions
}
