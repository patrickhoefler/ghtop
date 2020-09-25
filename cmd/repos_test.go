package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/mattn/go-shellwords"
	"github.com/patrickhoefler/ghtop/internal"
)

func Test_repos(t *testing.T) {
	tests := []struct {
		name          string
		cmd           string
		expectErr     bool
		expectResults int
	}{
		{
			name:          "default command",
			cmd:           "repos",
			expectResults: 100,
		},
		{
			name:          "min valid repo count",
			cmd:           "repos --fetch-repos 1",
			expectResults: 1,
		},
		{
			name:          "more than one page",
			cmd:           "repos --fetch-repos 123",
			expectResults: 123,
		},
		{
			name:          "max valid repo count",
			cmd:           "repos --fetch-repos 1000",
			expectResults: 1000,
		},
		{
			name:      "repo count not a number",
			cmd:       "repos --fetch-repos abc",
			expectErr: true,
		},
		{
			name:      "repo count negative",
			cmd:       "repos --fetch-repos -1",
			expectErr: true,
		},
		{
			name:      "repo count zero",
			cmd:       "repos --fetch-repos 0",
			expectErr: true,
		},
		{
			name:      "repo count too big",
			cmd:       "repos --fetch-repos 1001",
			expectErr: true,
		},
		{
			name:          "result limit",
			cmd:           "repos --top 2",
			expectResults: 2,
		},
		{
			name:      "result limit not a number",
			cmd:       "repos --top abc",
			expectErr: true,
		},
		{
			name:      "result limit negative",
			cmd:       "repos --top -1",
			expectErr: true,
		},
		{
			name:      "result limit zero",
			cmd:       "repos --top 0",
			expectErr: true,
		},
		{
			name:      "result limit too big",
			cmd:       "repos --top 1001",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			rootCmd := newRootCmd(internal.NewMockRepoSearchClient(), buf)

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

			output, err := ioutil.ReadAll(buf)
			if err != nil {
				panic(err)
			}

			lines := strings.Count(string(output), "\n")
			results := lines - 1

			if tt.expectResults > 0 && tt.expectResults != results {
				t.Fail()
				t.Logf("Expected %d results, but got %d", tt.expectResults, results)
				t.Logf("\n" + string(output))
			}
		})
	}
}
