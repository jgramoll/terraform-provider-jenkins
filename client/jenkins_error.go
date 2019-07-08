package client

import (
	"fmt"
)

// JenkinsError Error response from jenkins
type JenkinsError struct {
	ErrorMsg  string `xml:"error"`
	Exception string `xml:"exception"`
	Message   string `xml:"message"`
	Status    int    `xml:"status"`
	Timestamp int64  `xml:"timestamp"`
	Body      string `xml:"body"`
}

// For error interface
func (r *JenkinsError) Error() string {
	return fmt.Sprintf("%d %v: %v%v\n%v", r.Status, r.ErrorMsg, r.Message,
		r.Body, r.Exception)
}
