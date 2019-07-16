package client

type GitScmBranchSpec struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name"`
}

func NewGitScmBranchSpec() *GitScmBranchSpec {
	return &GitScmBranchSpec{}
}
