package repo

import (
	"context"
	"fmt"
)

func GetCurrentCommit(repository Repo) string {

	client := repository.getClient(repository)
	ctx := context.Background()

	branch, _, err := client.Repositories.GetBranch(ctx, repository.Owner, repository.Name, repository.Branch)

	if err != nil {
		fmt.Println(err)
	}

	return *branch.Commit.SHA
}
