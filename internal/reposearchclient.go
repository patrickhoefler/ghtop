package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
)

const maxReposPerPage = 100

func newRepoSearchClient(searchService repositorieser) Fetcher {
	return &repoSearchClient{search: searchService}
}

type repoSearchClient struct {
	search repositorieser
}

// Fetch fetches a list of repositories
func (rsc *repoSearchClient) Fetch(numberOfRepos int) []*github.Repository {
	opt := github.SearchOptions{
		Sort: "stars",
	}

	if numberOfRepos < maxReposPerPage {
		opt.ListOptions.PerPage = numberOfRepos
	} else {
		opt.ListOptions.PerPage = maxReposPerPage
	}

	var repos []*github.Repository

	for {
		fetchedRepos, resp, err := rsc.search.Repositories(
			context.Background(),
			"stars:>10000",
			&opt,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		repos = append(repos, fetchedRepos.Repositories...)

		if len(repos) >= numberOfRepos {
			break
		}

		opt.Page = resp.NextPage
	}

	return repos[:numberOfRepos]
}
