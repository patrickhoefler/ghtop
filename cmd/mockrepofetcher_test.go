package cmd

import "github.com/google/go-github/v32/github"

type mockRepoFetcher struct {
	topics []string
}

// Fetch always returns an empty []*github.Repository (for now)
// TODO: Clean up
func (mrf *mockRepoFetcher) Fetch(numberOfRepos int) (repos []*github.Repository) {
	repos = make([]*github.Repository, numberOfRepos)
	if len(mrf.topics) == 0 {
		mrf.topics = []string{"foo", "bar"}
	}

	stargazers1 := 123
	branch1 := "main"
	repos[0] = &github.Repository{
		Topics:          mrf.topics[:1],
		DefaultBranch:   &branch1,
		StargazersCount: &stargazers1,
	}

	if len(repos) > 1 {
		stargazers2 := 42
		branch2 := "dev"
		repos[1] = &github.Repository{
			Topics:          mrf.topics,
			DefaultBranch:   &branch2,
			StargazersCount: &stargazers2,
		}
	}

	return
}
