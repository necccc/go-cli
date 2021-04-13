package repo

import (
	"context"
	"fmt"

	"github.com/google/go-github/v34/github"
)

func GetLatestRelease(repository Repo) {

	client := repository.getClient(repository)
	ctx := context.Background()

	options := github.ListOptions{}

	tags, _, err := client.Repositories.ListTags(ctx, repository.Owner, repository.Name, &options)

	if err != nil {
		fmt.Println(err)
	}

	for _, tag := range tags {
		fmt.Println(*tag.Name)
	}

	// return tags[0]
}
