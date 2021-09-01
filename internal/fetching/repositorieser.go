package fetching

import (
	"context"

	"github.com/google/go-github/v38/github"
)

type repositorieser interface {
	Repositories(
		ctx context.Context,
		query string,
		opts *github.SearchOptions,
	) (
		*github.RepositoriesSearchResult,
		*github.Response,
		error,
	)
}
