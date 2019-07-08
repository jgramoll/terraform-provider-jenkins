package client

// Job
type Job struct {
	// actions
	Description       string `xml:"description"`
	DisplayName       string `xml:"displayName"`
	DisplayNameOrNull string `xml:"displayNameOrNull"`
	FullDisplayName   string `xml:"fullDisplayName"`
	FullName          string `xml:"fullName"`
	Name              string `xml:"name"`
	URL               string `xml:"url"`
	Buildable         bool   `xml:"buildable"`
	// builds
	Color string `xml:"color"`
	// firstBuild
	// healthReport
	// inQueue
	// keepDependencies
	// lastBuild
	// lastCompletedBuild
	// lastFailedBuild
	// lastStableBuild
	// lastSuccessfulBuild
	// lastUnstableBuild
	// lastUnsuccessfulBuild
	NextBuildNumber int64                 `xml:"nextBuildNumber"`
	Property        *[]*JobConfigProperty `xml:"property"`
	// queueItem
	ConcurrentBuild bool `xml:"concurrentBuild"`
	ResumeBlocked   bool `xml:"resumeBlocked"`
}

// NewJob return Job object with default values
func NewJob() *Job {
	return &Job{}
}
