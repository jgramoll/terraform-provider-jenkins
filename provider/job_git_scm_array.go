package provider

import (
	"github.com/jgramoll/terraform-provider-jenkins/client"
)

type jobGitScmArray []*jobGitScm

func (scmArray *jobGitScmArray) toClientSCM() (*client.GitSCM, error) {
	for _, scm := range *scmArray {
		return scm.toClientSCM()
	}
	return nil, nil
}

func (*jobGitScmArray) fromClientSCM(clientSCM *client.GitSCM) (*jobGitScmArray, error) {
	scm := newJobGitScm()
	clientScm, err := scm.fromClientSCM(clientSCM)
	if err != nil {
		return nil, err
	}
	return &jobGitScmArray{clientScm}, nil
}
