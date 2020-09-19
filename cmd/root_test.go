package cmd

import (
	"os"
	"testing"
)

func TestRootCmd(t *testing.T) {
	err := newRootCmd(new(mockRepoFetcher), os.Stdout).Execute()
	if err != nil {
		t.Fatalf("Could not run root command: %s", err)
	}
}
