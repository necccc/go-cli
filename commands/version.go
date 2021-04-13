package commands

import (
	"os"

	"github.com/necccc/go-cli/log"
	"github.com/necccc/go-cli/repo"
)

func Version(repository repo.Repo) {
	latestRelease := repo.GetLatestRelease(repository)

	log.Write("Latest release is %s \n", *latestRelease.Name)

	log.Write("Commit hash at release %s is %s \n", *latestRelease.Name, *latestRelease.Commit.SHA)

	currentCommit := repo.GetCurrentCommit(repository)

	log.Write("Current commit hash is %s \n", currentCommit)

	if currentCommit == *latestRelease.Commit.SHA {
		log.Write("No commits found since latest release\n")
		os.Exit(1)
	}

	log.Write("List of commits since %s: \n", *latestRelease.Name)


  
	log.Write("Next version is %s \n", "foo")

}
