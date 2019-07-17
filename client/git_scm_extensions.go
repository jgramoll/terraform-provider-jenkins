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
	var newExtensionItems []GitScmExtension
	if extensions.Items != nil {
		newExtensionItems = append(*extensions.Items, extension)
	} else {
		newExtensionItems = append(newExtensionItems, extension)
	}
	newExtensions := NewGitScmExtensions()
	newExtensions.Items = &newExtensionItems
	return newExtensions
}
