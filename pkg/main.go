package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/necccc/go-cli/commands"
	"github.com/necccc/go-cli/repo"
)

var appname string
var repoOwner string
var repoName string
var branch string
var versionPrefix string
var apiUrl string

func init() {

	flag.StringVar(&appname, "appname", "", "Application name, may contain [a-zA-Z0-9] and [_-], required")
	flag.StringVar(&repoOwner, "owner", "", "Repository owner, required")
	flag.StringVar(&repoName, "repo", "", "Repository name, required")

	flag.StringVar(&branch, "branch", "master", "Main branch to check PRs against")
	flag.StringVar(&versionPrefix, "version-prefix", "", "Version prefix, like 'v' for 'v2.3.4'. (default is empty)")
	flag.StringVar(&apiUrl, "api", "", "Your GitHub Enterprise API URL (default is empty)")
}

func main() {
	flag.Parse()

	if appname == "" || repoOwner == "" || repoName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	repository := repo.Repo{Owner: repoOwner, Name: repoName, ApiUrl: apiUrl}

	switch flag.Arg(0) {
	case "version":
		commands.Version(repository)
	case "release":
		fmt.Println("TBD: create release")
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}
