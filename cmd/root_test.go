package cmd

import (
	"os"
	"testing"

	"github.com/patrickhoefler/ghtop/internal/fetching"
)

func TestRootCmd(t *testing.T) {
	err := newRootCmd(fetching.NewMockRepoSearchClient(), os.Stdout).Execute()
	if err != nil {
		t.Fatalf("Could not run root command: %s", err)
	}
}
