package main

import (
	"fmt"
	"github.com/onedss/onedarwin/core"
	"github.com/onedss/onedarwin/routers"
	"log"
)

var (
	gitCommitCode string
	buildDateTime string
)

func main() {
	log.Printf("git commit code:%s", gitCommitCode)
	log.Printf("build date:%s", buildDateTime)
	routers.BuildVersion = fmt.Sprintf("%s.%s", routers.BuildVersion, gitCommitCode)
	routers.BuildDateTime = buildDateTime

	core.StartApp()
}
