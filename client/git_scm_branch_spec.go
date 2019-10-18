package client

type GitScmBranchSpec struct {
	Name string `xml:"name"`
}

func NewGitScmBranchSpec() *GitScmBranchSpec {
	return &GitScmBranchSpec{}
}
