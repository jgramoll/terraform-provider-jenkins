package client

type GitUserRemoteConfig struct {
	Id            string `xml:"id,attr,omitempty"`
	Refspec       string `xml:"refspec"`
	Url           string `xml:"url"`
	CredentialsId string `xml:"credentialsId"`
}

func NewGitUserRemoteConfig() *GitUserRemoteConfig {
	return &GitUserRemoteConfig{}
}
