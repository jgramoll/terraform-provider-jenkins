package client

type GitUserRemoteConfig struct {
	Refspec       string `xml:"refspec"`
	Url           string `xml:"url"`
	CredentialsId string `xml:"credentialsId"`
}

func NewGitUserRemoteConfig() *GitUserRemoteConfig {
	return &GitUserRemoteConfig{}
}
