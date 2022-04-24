package main

import "github.com/onedss/onedarwin/core"

var (
	gitCommitCode string
	buildDateTime string
)

func main() {
	core.SetGitCommitCode(gitCommitCode)
	core.SetBuildDateTime(buildDateTime)
	core.StartApp()
}
