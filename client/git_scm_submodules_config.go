package client

type GitScmSubmodulesConfig struct {
	Class string `xml:"class,attr"`
}

func NewGitScmSubmodulesConfig() *GitScmSubmodulesConfig {
	return &GitScmSubmodulesConfig{
		Class: "list",
	}
}
