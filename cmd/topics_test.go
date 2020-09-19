package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/mattn/go-shellwords"
)

func Test_topics(t *testing.T) {
	tests := []struct {
		name          string
		cmd           string
		expectErr     bool
		expectResults int
	}{
		{
			name:          "default command",
			cmd:           "topics",
			expectResults: 2,
		},
		{
			name:          "limit repos",
			cmd:           "topics --fetch-repos 10",
			expectResults: 2,
		},
		{
			name:          "limit results",
			cmd:           "topics --top 1",
			expectResults: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd := newRootCmd(&mockRepoFetcher{topics: []string{"foo", "bar"}}, buf)

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
