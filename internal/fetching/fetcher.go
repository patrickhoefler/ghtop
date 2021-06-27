package fetching

import "github.com/google/go-github/v36/github"

// Fetcher defines the interface for repo fetchers
type Fetcher interface {
	Fetch(n int) (repos []*github.Repository)
}
