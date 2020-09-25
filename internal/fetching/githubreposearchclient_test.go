package fetching

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v32/github"
)

func TestNewGitHubRepoSearchClient(t *testing.T) {
	tests := []struct {
		name string
		want Fetcher
	}{
		{
			name: "default",
			want: newRepoSearchClient(github.NewClient(nil).Search),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitHubRepoSearchClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitHubRepoSearchClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
