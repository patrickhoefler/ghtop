package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
)

const maxReposPerPage = 100

// RepoFetcher defines the interface for repo fetchers
type RepoFetcher interface {
	Fetch(n int) (repos []*github.Repository)
}

// GitHubRepoFetcher fetches the most starred repos
type GitHubRepoFetcher struct {
}

// Fetch bla bla bla
func (ghrf *GitHubRepoFetcher) Fetch(numberOfRepos int) (repos []*github.Repository) {
	client := github.NewClient(nil)
	opt := github.SearchOptions{
		Sort: "stars",
	}

	pages := ((numberOfRepos - 1) / maxReposPerPage) + 1
	for pageCounter := 0; pageCounter < pages; pageCounter++ {
		remainingRepos := numberOfRepos - (pageCounter * maxReposPerPage)
		if remainingRepos >= maxReposPerPage {
			opt.ListOptions.PerPage = maxReposPerPage
		} else {
			opt.ListOptions.PerPage = remainingRepos
		}

		fetchedRepos, resp, err := client.Search.Repositories(
			context.Background(),
			"stars:>10000",
			&opt,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		repos = append(repos, fetchedRepos.Repositories...)
		opt.Page = resp.NextPage
	}

	return
}
