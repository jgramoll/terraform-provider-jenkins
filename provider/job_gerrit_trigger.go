package provider

type skipVote struct {
	OnSuccessful bool `xml:"on_successful"`
	OnFailed bool `xml:"on_failed"`
	OnUnstable bool `xml:"on_unstable"`
	OnNotBuilt bool `xml:"on_not_built"`
}

func newSkipVote() *skipVote {
	return &skipVote{}
}

type jobGerritTrigger struct {
	Property string `xml:"property"`
	ServerName string `xml:"server_name"`
	SilentMode bool `xml:"silent_mode"`
	SilentStartMode bool `xml:"silent_start_mode"`
	EscapeQuotes bool `xml:"escape_quotes"`
	NameAndEmailParameterMode string `xml:"name_and_email_parameter_mode"`
	CommitMessageParameterMode string `xml:"commit_message_parameter_mode"`
	ChangeSubjectParameterMode string `xml:"change_subject_parameter_mode"`
	CommentTextParameterMode string `xml:"comment_text_parameter_mode"`

	SkipVote *skipVote `xml:"skip_vote"`
}

func newJobGerritTrigger() *jobGerritTrigger {
	return &jobGerritTrigger{
		EscapeQuotes: true,
		SkipVote: newSkipVote(),
	}
}
