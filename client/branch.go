package client

type Branches struct {
	Items *[]*Branch `xml:"hudson.plugins.git.BranchSpec"`
}

type Branch struct {
	Name string `xml:"name"`
}
