package fetching

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-github/v33/github"
)

func TestNewMockRepoSearchClient(t *testing.T) {
	tests := []struct {
		name string
		want Fetcher
	}{
		{
			name: "default",
			want: newRepoSearchClient(new(mockGitHubSearchClient)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockRepoSearchClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockRepoSearchClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mockGitHubSearchClient_Repositories(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
		opts  *github.SearchOptions
	}
	tests := []struct {
		name             string
		mghcs            *mockGitHubSearchClient
		args             args
		wantSearchResult *github.RepositoriesSearchResult
		wantResp         *github.Response
		wantErr          bool
	}{
		{
			name: "default",
			args: args{
				opts: &github.SearchOptions{
					ListOptions: github.ListOptions{
						PerPage: 100,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mghcs := &mockGitHubSearchClient{}
			_, _, err := mghcs.Repositories(tt.args.ctx, tt.args.query, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockGitHubSearchClient.Repositories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
