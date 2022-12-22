package fetching

import (
	"context"

	"github.com/google/go-github/v48/github"
)

// NewMockRepoSearchClient returns a new mock repo fetcher for testing
func NewMockRepoSearchClient() Fetcher {
	return newRepoSearchClient(new(mockGitHubSearchClient))
}

type mockGitHubSearchClient struct{}

func (
	mghcs *mockGitHubSearchClient,
) Repositories(
	ctx context.Context,
	query string,
	opts *github.SearchOptions,
) (
	searchResult *github.RepositoriesSearchResult,
	resp *github.Response,
	err error,
) {
	var repos []*github.Repository

	if opts.Page == 0 {
		opts.Page = 1
	}

	for repoCounter := 0; repoCounter < opts.PerPage; repoCounter++ {
		stargazerCount := new(int)
		*stargazerCount = 10000000 - (opts.Page-1)*opts.PerPage - repoCounter
		defaultBranch := new(string)
		*defaultBranch = "main"

		if *stargazerCount == 10000000 {
			repos = append(repos, &github.Repository{
				DefaultBranch:   defaultBranch,
				StargazersCount: stargazerCount,
				Topics:          []string{"foo"},
			})
		} else {
			repos = append(repos, &github.Repository{
				StargazersCount: stargazerCount,
				Topics:          []string{"foo", "bar"},
			})
		}
	}

	searchResult = &github.RepositoriesSearchResult{
		Repositories: repos,
	}

	resp = &github.Response{
		NextPage: opts.Page + 1,
	}

	return
}
