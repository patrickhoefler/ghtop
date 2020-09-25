package cmd

import (
	"os"
	"testing"

	"github.com/patrickhoefler/ghtop/internal"
)

func TestRootCmd(t *testing.T) {
	err := newRootCmd(internal.NewMockRepoSearchClient(), os.Stdout).Execute()
	if err != nil {
		t.Fatalf("Could not run root command: %s", err)
	}
}
