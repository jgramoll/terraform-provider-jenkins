package client

type GitScmExtensions struct {
	Items *[]GitScmExtension
}

func NewGitScmExtensions() *GitScmExtensions {
	return &GitScmExtensions{
		Items: &[]GitScmExtension{},
	}
}

func (extensions *GitScmExtensions) Append(extension GitScmExtension) *GitScmExtensions {
	newExtensions := NewGitScmExtensions()
	if extensions.Items != nil {
		*newExtensions.Items = append(*extensions.Items, extension)
	} else {
		*newExtensions.Items = []GitScmExtension{extension}
	}
	return newExtensions
}
