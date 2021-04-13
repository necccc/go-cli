package repo

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

type Repo struct {
	Owner, Name, ApiUrl, Branch, VersionPrefix string
}

func createOauthClient() *http.Client {
	token := os.Getenv("GIT_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return tc
}

func (r Repo) getClient(repository Repo) *github.Client {
	tc := createOauthClient()

	var client *github.Client
	var err error

	if repository.ApiUrl == "" {
		client = github.NewClient(tc)
	} else {
		client, err = github.NewEnterpriseClient(repository.ApiUrl, repository.ApiUrl, tc)

		if err != nil {
			fmt.Println(err)
		}
	}

	return client
}
