package fetching

import (
	"github.com/google/go-github/v39/github"
)

// NewGitHubRepoSearchClient returns a new GitHub repo fetcher
func NewGitHubRepoSearchClient() Fetcher {
	return newRepoSearchClient(github.NewClient(nil).Search)
}
