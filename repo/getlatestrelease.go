package repo

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/go-github/v34/github"
)

func GetLatestRelease(repository Repo) *github.RepositoryTag {

	client := repository.getClient(repository)
	ctx := context.Background()

	options := github.ListOptions{}

	tags, _, err := client.Repositories.ListTags(ctx, repository.Owner, repository.Name, &options)

	if err != nil {
		fmt.Println(err)
	}

	tags = FilterReleases(tags, repository)
	releases := OrderReleases(tags, repository)

	return releases[0]
}

func FilterReleases(tags []*github.RepositoryTag, repository Repo) []*github.RepositoryTag {
	releases := make([]*github.RepositoryTag, 0)

	prefix := repository.VersionPrefix
	re := regexp.MustCompile(`^` + regexp.QuoteMeta(prefix) + `\d*\.\d*.\d*$`)

	fmt.Println(re)

	for _, tag := range tags {
		if re.MatchString(*tag.Name) {
			releases = append(releases, tag)
		}
	}

	return releases
}

func toInt(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}

func OrderReleases(tags []*github.RepositoryTag, repository Repo) []*github.RepositoryTag {
	sort.Slice(tags, func(i, j int) bool {
		next := strings.Split(*tags[i].Name, ".")
		prev := strings.Split(*tags[j].Name, ".")

		if toInt(next[0]) != toInt(prev[0]) {
			return toInt(next[0]) > toInt(prev[0])
		}

		if toInt(next[1]) != toInt(prev[1]) {
			return toInt(next[1]) > toInt(prev[1])
		}

		return toInt(next[2]) > toInt(prev[2])
	})

	return tags
}
