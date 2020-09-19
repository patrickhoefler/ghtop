package cmd

import (
	"bytes"
	"log"
	"testing"

	"github.com/mattn/go-shellwords"
)

func Test_defaultbranches(t *testing.T) {
	tests := []struct {
		name      string
		cmd       string
		expectErr bool
	}{
		{
			name: "default command",
			cmd:  "defaultbranches",
		},
		{
			name: "limit repos",
			cmd:  "defaultbranches --fetch-repos 10",
		},
		{
			name: "limit results",
			cmd:  "defaultbranches --top 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd := newRootCmd(new(mockRepoFetcher), buf)

			args, err := shellwords.Parse(tt.cmd)
			if err != nil {
				log.Fatal("Could not parse command args")
			}
			rootCmd.SetArgs(args)

			err = rootCmd.Execute()
			if tt.expectErr && err == nil {
				t.Fail()
				t.Logf("Exptected an error, but non occurred")
			} else if !tt.expectErr && err != nil {
				t.Fail()
				t.Logf("Expected no error, but got: %s", err)
			}
		})
	}
}
