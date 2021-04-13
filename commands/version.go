package commands

import "github.com/necccc/go-cli/repo"

func Version(repository repo.Repo) {
	repo.GetLatestRelease(repository)
}
