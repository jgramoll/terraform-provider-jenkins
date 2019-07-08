package client

type UserRemoteConfigs struct {
	Items *[]*UserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}

type UserRemoteConfig struct {
	Url           string `xml:"url"`
	CredentialsId string `xml:"credentialsId"`
}
