package fetching

import (
	"testing"

	"github.com/google/go-github/v48/github"
)

func Test_newRepoSearchClient(t *testing.T) {
	type args struct {
		searchService repositorieser
	}
	tests := []struct {
		name string
		args args
		want Fetcher
	}{
		{
			name: "GitHub",
			args: args{searchService: &github.SearchService{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRepoSearchClient(tt.args.searchService)
		})
	}
}

func Test_repoSearchClient_Fetch(t *testing.T) {
	type fields struct {
		search repositorieser
	}
	type args struct {
		numberOfRepos int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults int
	}{
		{
			name:        "default",
			fields:      fields{search: &mockGitHubSearchClient{}},
			args:        args{numberOfRepos: 100},
			wantResults: 100,
		},
		{
			name:        "fewer than default page size",
			fields:      fields{search: &mockGitHubSearchClient{}},
			args:        args{numberOfRepos: 42},
			wantResults: 42,
		},
		{
			name:        "larger than default page size",
			fields:      fields{search: &mockGitHubSearchClient{}},
			args:        args{numberOfRepos: 123},
			wantResults: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsc := &repoSearchClient{
				search: tt.fields.search,
			}

			results := len(rsc.Fetch(tt.args.numberOfRepos))
			if tt.wantResults > 0 && tt.wantResults != results {
				t.Fail()
				t.Logf("Wanted %d results, but got %d", tt.wantResults, results)
			}
		})
	}
}
